package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/example/internal/entity"
	"github.com/example/internal/repository"
	"github.com/example/internal/utils"
)

type payrollService struct {
	payrollRepo      repository.Repository[entity.Payroll]
	clockRepo        repository.Repository[entity.ClockEntry]
	leaveRepo        repository.Repository[entity.LeaveRequest]
	employeeRepo     repository.Repository[entity.Employee]
	leaveBalanceRepo repository.Repository[entity.LeaveBalance]
}

type PayrollService interface {
	GeneratePayroll(ctx context.Context, employeeID string, month time.Month, year int) (any, error)
}

func NewPayrollService(
	payrollRepo repository.Repository[entity.Payroll],
	clockRepo repository.Repository[entity.ClockEntry],
	leaveRepo repository.Repository[entity.LeaveRequest],
	employeeRepo repository.Repository[entity.Employee],
	leaveBalanceRepo repository.Repository[entity.LeaveBalance],
) PayrollService {
	return &payrollService{
		payrollRepo:      payrollRepo,
		clockRepo:        clockRepo,
		leaveRepo:        leaveRepo,
		employeeRepo:     employeeRepo,
		leaveBalanceRepo: leaveBalanceRepo,
	}
}

// helper function
func (s *payrollService) GeneratePayroll(ctx context.Context, employeeID string, month time.Month, year int) (any, error) {
	fmt.Print(employeeID)
	employees, err := s.employeeRepo.Query(ctx, map[string]interface{}{
		"id": employeeID,
	})
	if err != nil {
		// fmt.Print(err)
		return nil, err
	}
	if len(employees) == 0 {
		return nil, errors.New("employee not found")
	}
	employee := employees[0]
	location := time.UTC
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, location)
	endDate := startDate.AddDate(0, 1, -1)

	// If LWD is set and it's before endDate, cap the payroll period
	if !employee.EmploymentEndDate.IsZero() && employee.EmploymentEndDate.Before(endDate) {
		endDate = employee.EmploymentEndDate
	}

	// Get clock entries
	clockEntries, err := s.clockRepo.Query(ctx, map[string]interface{}{
		"employee_id": employeeID,
	})
	if err != nil {
		return nil, err
	}

	// Get approved leaves
	leaveRequests, err := s.leaveRepo.Query(ctx, map[string]interface{}{
		"employee_id": employeeID,
		"status":      "APPROVED",
	})
	if err != nil {
		return nil, err
	}

	// Get leave balance
	balances, err := s.leaveBalanceRepo.Query(ctx, map[string]interface{}{
		"employee_id": employeeID,
		"month":       int(month),
		"year":        year,
	})
	if err != nil || len(balances) == 0 {
		return nil, errors.New("leave balance not found")
	}
	leaveBalance := balances[0]

	clockedDates := make(map[string]float64)
	for _, c := range clockEntries {
		date := c.ClockInTime.In(location).Format("2006-01-02")
		duration := c.ClockOutTime.Sub(c.ClockInTime).Hours()
		clockedDates[date] += duration
	}

	leaveDates := make(map[string]entity.LeaveRequest)
	for _, lr := range leaveRequests {
		for d := lr.StartDate; !d.After(lr.EndDate); d = d.AddDate(0, 0, 1) {
			leaveDates[d.Format("2006-01-02")] = lr
		}
	}

	totalHours := 0.0
	overTime := 0.0
	paidLeaves := 0
	unpaidLeaves := 0

	hoursPerDay := 8.0

	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			continue
		}
		key := d.Format("2006-01-02")

		if hrs, ok := clockedDates[key]; ok {
			totalHours += hrs
			if hrs > hoursPerDay {
				overTime += hrs - hoursPerDay
			}
		} else if leave, ok := leaveDates[key]; ok {
			if entity.LeaveType(leave.LeaveType) == entity.UnpaidLeave || leaveBalance.RemPaid == 0 {
				unpaidLeaves++
			} else {
				paidLeaves++
				leaveBalance.RemPaid--
			}
		} else {
			unpaidLeaves++
		}
	}

	deductions := float64(unpaidLeaves) * (employee.BaseSalary / 30)
	finalSalary := employee.BaseSalary + employee.OvertimeRate*overTime - deductions

	payroll := entity.Payroll{
		Month:         startDate,
		EmployeeId:    employeeID,
		BaseSalary:    employee.BaseSalary,
		TotalHours:    totalHours,
		OverTimeHours: overTime,
		PaidLeaves:    float64(paidLeaves),
		UnpaidLeaves:  float64(unpaidLeaves),
		Deductions:    deductions,
		Bonuses:       0,
		FinalSalary:   finalSalary,
		PdfUrl:        "",
		Status:        entity.Pending,
		PaymentDate:   time.Time{},
	}

	pdfBuf, err := utils.GeneratePayrollPDF(payroll, employee.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate pdf: %w", err)
	}

	s3Url, err := utils.UploadToS3(ctx, pdfBuf, fmt.Sprintf("%s-%d-%d.pdf", employeeID, month, year))
	if err != nil {
		return nil, fmt.Errorf("failed to upload pdf: %w", err)
	}

	payroll.PdfUrl = s3Url
	_ = s.payrollRepo.Update(ctx, payroll.ID, map[string]interface{}{"pdf_url": s3Url})

	// Optional: Update leave balance (persistence)
	err = s.leaveBalanceRepo.Update(ctx, leaveBalance.ID, map[string]interface{}{
		"rem_paid": leaveBalance.RemPaid,
	})
	if err != nil {
		return nil, err
	}

	return s.payrollRepo.Create(ctx, &payroll)
}

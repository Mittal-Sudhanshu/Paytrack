package service

import (
	"context"
	"fmt"
	"time"

	"github.com/example/internal/entity"
	"github.com/example/internal/model"
	"github.com/example/internal/repository"
)

type leaveService struct {
	leaveRequestRepo repository.Repository[entity.LeaveRequest]
	leaveBalanceRepo repository.Repository[entity.LeaveBalance]
	employeeRepo     repository.Repository[entity.Employee]
}

type LeaveService interface {
	ApplyLeave(context context.Context, userId string, leaveRequest model.LeaveRequest) (any, error)
	GetLeaveBalance(context context.Context, userId string) (any, error)
	// GetLeaveHistory(context context.Context, userId string) ([]entity.LeaveRequest, error)
	GetLeaveRequests(context context.Context, userId string) ([]entity.LeaveRequest, error)
	UpdateLeaveRequest(context context.Context, requestId string, userId string, status string) (any, error)
	GenerateLeaveBalance(ctx context.Context) (any, error)
}

func NewLeaveService(leaveRequestRepo repository.Repository[entity.LeaveRequest], leaveBalanceRepo repository.Repository[entity.LeaveBalance], employeeRepo repository.Repository[entity.Employee]) LeaveService {
	return &leaveService{
		leaveRequestRepo: leaveRequestRepo,
		leaveBalanceRepo: leaveBalanceRepo,
		employeeRepo:     employeeRepo,
	}
}

func (s *leaveService) ApplyLeave(context context.Context, userId string, leaveRequest model.LeaveRequest) (any, error) {
	// Implement the logic to apply leave
	leaveRequestEntity := entity.LeaveRequest{
		EmployeeId: userId,
		LeaveType:  leaveRequest.LeaveType,
		StartDate:  leaveRequest.StartDate,
		EndDate:    leaveRequest.EndDate,
		Reason:     leaveRequest.Reason,
	}
	//save to db
	leave, err := s.leaveRequestRepo.Create(context, &leaveRequestEntity)
	if err != nil {
		return nil, err
	}
	return leave, nil
}

func (s *leaveService) GetLeaveBalance(context context.Context, userId string) (any, error) {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	leaveBalance, err := s.leaveBalanceRepo.Query(context, map[string]interface{}{
		"employee_id": userId,
		"year":        currentYear,
		"month":       currentMonth,
	})
	if err != nil {
		return nil, err
	}
	return leaveBalance[0], nil
}

func (s *leaveService) GetLeaveRequests(context context.Context, userId string) ([]entity.LeaveRequest, error) {
	leaveRequests, err := s.leaveRequestRepo.Query(context, map[string]interface{}{
		"employee_id": userId,
	})
	if err != nil {
		return nil, err
	}
	return leaveRequests, nil
}
func (s *leaveService) UpdateLeaveRequest(context context.Context, requestId string, userId string, status string) (any, error) {
	// Fetch the leave request
	leaveReq, err := s.leaveRequestRepo.GetByID(context, requestId)
	if err != nil {
		return nil, err
	}

	// If not approved, just update status and return
	if status != "APPROVED" {
		return s.leaveRequestRepo.Update(context, requestId, map[string]interface{}{"status": status}), nil
	}

	// Calculate number of leave days
	numDays := int(leaveReq.EndDate.Sub(leaveReq.StartDate).Hours()/24) + 1

	// Fetch leave balance
	balanceList, err := s.leaveBalanceRepo.Query(context, map[string]interface{}{
		"employee_id": leaveReq.EmployeeId,
	})
	if err != nil || len(balanceList) == 0 {
		return nil, err
	}
	balance := balanceList[0]

	// Handle balance deduction
	switch leaveReq.LeaveType {
	case 0, 1:
		if balance.RemPaid >= numDays {
			balance.RemUnpaid -= numDays
		} else {
			requiredUnpaid := numDays - balance.RemPaid
			if balance.RemUnpaid < requiredUnpaid {
				return nil, fmt.Errorf("not enough leave balance: paid and unpaid combined")
			}
			balance.RemUnpaid -= requiredUnpaid
			balance.RemPaid = 0
		}
	case 5:
		if balance.RemUnpaid < numDays {
			return nil, fmt.Errorf("not enough unpaid leave balance")
		}
		balance.RemUnpaid -= numDays
	}

	// Save updated balance
	err = s.leaveBalanceRepo.Update(context, balance.ID, map[string]interface{}{
		"rem_paid":   balance.RemPaid,
		"rem_unpaid": balance.RemUnpaid,
	})
	if err != nil {
		return nil, err
	}

	// Finally, update the leave request status
	err = s.leaveRequestRepo.Update(context, requestId, map[string]interface{}{
		"status": status,
	})
	if err != nil {
		return nil, err
	}
	return "Leave Approved", nil
}

func (s *leaveService) GenerateLeaveBalance(ctx context.Context) (any, error) {
	now := time.Now()
	month := int(now.Month())
	year := now.Year()

	// 1. Fetch all active employees
	employees, err := s.employeeRepo.Query(ctx, map[string]interface{}{
		"is_active": false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employees: %w", err)
	}

	// 2. For each employee, check if balance already exists for the month
	for _, e := range employees {
		existing, err := s.leaveBalanceRepo.Query(ctx, map[string]interface{}{
			"employee_id": e.ID,
			"month":       int(month),
			"year":        year,
		})
		if err != nil {
			return nil, fmt.Errorf("error querying leave balance: %w", err)
		}
		if len(existing) > 0 {
			continue // Skip if already generated
		}

		// 3. Create new leave balance (customize totals as needed)
		leaveBalance := entity.LeaveBalance{
			EmployeeId:  e.ID,
			Month:       month,
			Year:        year,
			TotalPaid:   2, // Monthly entitlement, example value
			TotalUnpaid: 10,
			RemPaid:     2,
			RemUnpaid:   10,
		}

		if _, err := s.leaveBalanceRepo.Create(ctx, &leaveBalance); err != nil {
			return nil, fmt.Errorf("failed to create leave balance for %s: %w", e.ID, err)
		}
	}

	return "Success", nil
}

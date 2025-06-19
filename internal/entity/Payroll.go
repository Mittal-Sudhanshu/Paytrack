package entity

import "time"

type PayrollStatus int

const (
	Paid    PayrollStatus = 1
	Pending PayrollStatus = 2
)

type Payroll struct {
	BaseModel
	Month         time.Time
	EmployeeId    string
	Employee      Employee `gorm:"foreignKey:EmployeeId"`
	BaseSalary    float64
	TotalHours    float64
	OverTimeHours float64
	PaidLeaves    float64
	UnpaidLeaves  float64
	Deductions    float64
	Bonuses       float64
	FinalSalary   float64
	PdfUrl        string
	Status        PayrollStatus
	PaymentDate   time.Time
}

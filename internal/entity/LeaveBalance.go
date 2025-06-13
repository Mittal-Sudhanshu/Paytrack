package entity

import "time"

type LeaveBalance struct {
	BaseModel
	EmployeeId  string
	Employee    Employee `gorm:"foreignKey:EmployeeId"`
	Month       time.Month
	Year        int
	TotalPaid   int
	TotalUnpaid int
	RemPaid     int
	RemUnpaid   int
}

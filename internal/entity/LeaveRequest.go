package entity

import "time"

type LeaveType int

const (
	CasualLeave    LeaveType = 0
	AnnualLeave    LeaveType = 1
	SickLeave      LeaveType = 2
	MaternityLeave LeaveType = 3
	PaternityLeave LeaveType = 4
	UnpaidLeave    LeaveType = 5
	OtherLeave     LeaveType = 6
)

type LeaveRequest struct {
	BaseModel
	EmployeeId string
	Employee   Employee `gorm:"foreignKey:EmployeeId"`
	StartDate  time.Time
	EndDate    time.Time
	LeaveType  int    `gorm:"type:integer;not null"`
	Status     string `gorm:"type:varchar(20);not null;default:'PENDING'"`
	Reason     string `gorm:"type:text"`
	ApprovedBy string
	Approver   Employee `gorm:"foreignKey:ApprovedBy"`
}

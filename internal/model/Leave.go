package model

import "time"

type LeaveRequest struct {
	LeaveType int       `json:"leaveType" binding:"required"`
	StartDate time.Time `json:"startDate" binding:"required"` // Format: YYYY-MM-DD
	EndDate   time.Time `json:"endDate" binding:"required"`   // Format: YYYY-MM-DD
	Reason    string       `json:"reason" binding:"required"`
}

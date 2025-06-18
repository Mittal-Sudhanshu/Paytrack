package entity

type LeaveBalance struct {
	BaseModel
	EmployeeId  string
	Employee    Employee `gorm:"foreignKey:EmployeeId"`
	Month       int      `gorm:"type:integer"` // ✅ Changed from time.Month to int
	Year        int
	TotalPaid   int
	TotalUnpaid int
	RemPaid     int
	RemUnpaid   int
}

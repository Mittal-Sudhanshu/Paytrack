package entity

import "time"

type ClockEntry struct {
	BaseModel
	EmployeeId   string
	Employee     Employee `gorm:"foreignKey:EmployeeId"`
	ClockInTime  time.Time
	ClockOutTime time.Time
	Timezone     string
}

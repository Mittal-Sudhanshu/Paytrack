package entity

import "time"

type Invite struct {
	BaseModel
	Email                   string       `gorm:"unique;not null"`
	RoleId                  string       `gorm:"not null"`
	Role                    Role         `gorm:"foreignKey:RoleId"`
	FirstName               string       `gorm:"not null, default:''"`
	LastName                string       `gorm:"not null, default:''"`
	InvitedById             string       `gorm:"not null"`
	InvitedBy               User         `gorm:"foreignKey:InvitedById"`
	OrganizationId          string       `gorm:"not null"`
	Organization            Organization `gorm:"foreignKey:OrganizationId"`
	ExpiresAt               time.Time    `gorm:"not null"`
	InviteToken             string       `gorm:"unique;not null"`
	Status                  InviteStatus `gorm:"not null,default:0"`
	Message                 string       `gorm:"type:text, default:'You are invited to join our organization.'"`
	Department              string
	Designation             string
	BaseSalary              float64
	Bonus                   float64
	OvertimeRate            float64
	Allowances              float64
	HealthInsurance         float64
	RetirementBenefits      float64
	StockOptions            float64
	StockOptionsVested      float64
	StockOptionsUnvested    float64
	StockOptionsStrikePrice float64
	StockOptionsQuantity    int
	StockOptionsType        string // e.g., "ISO", "NSO"
	StockOptionsStatus      string
	JoiningDate             time.Time
	EmploymentType          string  // e.g., "Full-time", "Part-time", "Contract"
	ReportingToID           *string // Optional, can be null if no direct manager
	PhoneNumber             string  `gorm:"not null, default:''"`
}

type InviteStatus int

const (
	InvitePending  InviteStatus = 0 // 0
	InviteAccepted InviteStatus = 1 // 1
	InviteDeclined InviteStatus = 2 // 2
)

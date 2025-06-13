package entity

import "time"

type Employee struct {
	BaseModel
	UserID         string `gorm:"not null"`
	OrganizationID string `gorm:"not null"`

	// Relationships with foreign keys
	User                    User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Organization            Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Designation             string
	FirstName               string
	LastName                string
	Department              string
	BaseSalary              float64
	Bonus                   float64
	OvertimeRate            float64
	Allowances              float64
	HealthInsurance         float64
	RetirementBenefits      float64
	StockOptions            float64
	StockOptionsVested      float64
	StockOptionsUnvested    float64
	StockOptionsExpiration  time.Time
	StockOptionsGrantDate   time.Time
	StockOptionsStrikePrice float64
	StockOptionsQuantity    int
	StockOptionsType        string // e.g., "ISO", "NSO"
	StockOptionsStatus      string // e.g., "Active", "Vested", "Expired"
	StockOptionsPlanID      string // Foreign key to Stock Options Plan, if applicable
	JoiningDate             time.Time
	ReportingToID           *string
	EmploymentType          string
	EmploymentEndDate       time.Time
}

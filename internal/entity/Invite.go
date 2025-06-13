package entity

import "time"

type Invite struct {
	BaseModel
	Email                   string       `gorm:"unique;not null"`
	RoleId                  string       `gorm:"not null"`
	Role                    Role         `gorm:"foreignKey:RoleId"`
	FirstName               string       `gorm:"not null;default:''"`
	LastName                string       `gorm:"not null;default:''"`
	InvitedById             string       `gorm:"not null"`
	InvitedBy               User         `gorm:"foreignKey:InvitedById"`
	OrganizationId          string       `gorm:"not null"`
	Organization            Organization `gorm:"foreignKey:OrganizationId"`
	ExpiresAt               time.Time    `gorm:"not null"`
	InviteToken             string       `gorm:"unique;not null"`
	Status                  InviteStatus `gorm:"not null;default:0"`
	Message                 string       `gorm:"type:text;default:'You are invited to join our organization.'"`
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
	StockOptionsType        string
	StockOptionsStatus      string
	JoiningDate             time.Time
	EmploymentType          string
	ReportingToID           *string
	PhoneNumber             string `gorm:"not null;default:''"`
}

type InviteStatus int

const (
	InvitePending  InviteStatus = 0 // 0
	InviteAccepted InviteStatus = 1 // 1
	InviteDeclined InviteStatus = 2 // 2
)

// CREATE TABLE "invites" ("id" uuid DEFAULT gen_random_uuid(),"created_at" timestamptz,"updated_at" timestamptz,"is_active" boolean,"deleted_at" timestamptz,"email" text NOT NULL,"role_id" uuid NOT NULL,"first_name" text,"last_name" text,"invited_by_id" uuid NOT NULL,"organization_id" uuid NOT NULL,"expires_at" timestamptz NOT NULL,"invite_token" text NOT NULL,"status" bigint,"message" text, default:'You are invited to join our organization.',"department" text,"designation" text,"base_salary" decimal,"bonus" decimal,"overtime_rate" decimal,"allowances" decimal,"health_insurance" decimal,"retirement_benefits" decimal,"stock_options" decimal,"stock_options_vested" decimal,"stock_options_unvested" decimal,"stock_options_strike_price" decimal,"stock_options_quantity" bigint,"stock_options_type" text,"stock_options_status" text,"joining_date" timestamptz,"employment_type" text,"reporting_to_id" text,"phone_number" text,PRIMARY KEY ("id"),CONSTRAINT "fk_invites_role" FOREIGN KEY ("role_id") REFERENCES "roles"("id"),CONSTRAINT "fk_invites_invited_by" FOREIGN KEY ("invited_by_id") REFERENCES "users"("id"),CONSTRAINT "fk_invites_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations"("id"),CONSTRAINT "uni_invites_email" UNIQUE ("email"),CONSTRAINT "uni_invites_invite_token" UNIQUE ("invite_token"))
// 2025/06/13 23:06:44 Failed to connect to DB: auto migration failed: ERROR: syntax error at or near "default" (SQLSTATE 42601)
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x102846c8c]

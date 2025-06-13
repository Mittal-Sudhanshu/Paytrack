package entity

import "time"

type User struct {
	BaseModel
	Email         string `gorm:"unique"`
	Password_hash string
	RoleId        string
	Role          Role `gorm:"forgienKey:RoleId"`
	FirstName     string
	LastName      string
	PhoneNumnber  string
	LastLogin     time.Time
}

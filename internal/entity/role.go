package entity

type Role struct {
	BaseModel
	Name        string
	Permissions []string `gorm:"type:json"`
	Scope       string
}

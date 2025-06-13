package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	IsActive  bool
	DeletedAt gorm.DeletedAt `gorm:"index"` // optional for soft deletes
}

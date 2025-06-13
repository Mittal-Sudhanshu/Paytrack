package entity

type UserOrg struct {
	BaseModel
	UserID string       `json:"user_id"`
	User   User         `gorm:"foreignKey:UserID"`
	OrgID  string       `json:"org_id"`
	Org    Organization `gorm:"foreignKey:OrgID"`
	RoleId string       `json:"role_id"`
	Role   Role         `gorm:"foreignKey:RoleId"`
}

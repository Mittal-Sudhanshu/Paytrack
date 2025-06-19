package model

type SignupRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	RoleId      string `json:"roleId" binding:"required"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

package model

type AcceptInviteRequest struct {
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	InviteToken string `json:"inviteToken" binding:"required"`
}

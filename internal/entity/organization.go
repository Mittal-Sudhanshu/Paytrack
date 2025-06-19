package entity

type Organization struct {
	BaseModel
	Name          string
	Industry      string
	Address       string
	City          string
	Country       string
	Website       string
	Contact_email string
	Contact_phone string
}

package models

type Customer struct {
	Customer_ID string `gorm:"primaryKey" json:"CustomerId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

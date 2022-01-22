package models

type User struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	PhoneNumber string `json:"phone_number"`
}

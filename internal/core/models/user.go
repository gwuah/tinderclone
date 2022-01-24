package models

import "gorm.io/datatypes"

type User struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	PhoneNumber string         `json:"phone_number"`
	OTP         string         `json:"otp"`
	Created     datatypes.Date `json:"created_at"`
}

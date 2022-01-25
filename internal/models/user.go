package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	PhoneNumber string    `json:"phone_number"`
	OTP         string    `json:"otp"`
	CreatedAt   time.Time `json:"created_at"`
}

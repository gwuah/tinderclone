package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id" gorm:"default:gen_random_uuid()"`
	CountryCode  string    `json:"country_code"`
	PhoneNumber  string    `json:"phone_number"`
	OTP          string    `json:"otp"`
	CreatedAt    time.Time `json:"created_at" sql:"type:timestamp without time zone" `
	OTPCreatedAt time.Time `json:"otp_created_at" sql:"type:timestamp without time zone" `
	FirstName    string    `json:"first_name"`
	DOB          time.Time `json:"dob" sql:"type:timestamp without time zone"`
	Location     string    `json:"location"`
}	
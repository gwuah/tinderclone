package models

import (
	"time"
)

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Scores struct {
	FirstName    int `json:"firstname"`
	LastName     int `json:"lastname"`
	Location     int `json:"location"`
	Bio          int `json:"bio"`
	Gender       int `json:"gender"`
	DOB          int `json:"dob"`
	Interests    int `json:"interests"`
	ProfilePhoto int `json:"profile_photo"`
}

type User struct {
	ID           string    `json:"id" gorm:"default:gen_random_uuid()"`
	CountryCode  string    `json:"country_code"`
	PhoneNumber  string    `json:"phone_number"`
	OTP          string    `json:"otp"`
	RawOTP       string    `json:"raw_otp"`
	CreatedAt    time.Time `json:"created_at" sql:"type:timestamp without time zone" `
	OTPCreatedAt time.Time `json:"otp_created_at" sql:"type:timestamp without time zone" `
	FirstName    string    `json:"first_name"`
	DOB          time.Time `json:"dob" sql:"type:timestamp without time zone"`
	Scores       Scores    `json:"scores" gorm:"-"`
}

func (u *User) Sanitize() {
	u.OTP = ""
	u.OTPCreatedAt = time.Time{}
}

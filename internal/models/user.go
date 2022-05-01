package models

import (
	"context"
	"fmt"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
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
	LastName     string    `json:"last_name"`
	Bio          string    `json:"bio"`
	Location     Location  `json:"location"`
	DOB          time.Time `json:"dob" sql:"type:timestamp without time zone"`
	Gender       string    `json:"gender"`
	Interests    string    `json:"interests"`
	Scores       Scores    `json:"scores" gorm:"-"`
	ProfilePhoto string    `json:"profile_photo"`
}

func (u *User) Sanitize() {
	u.OTP = ""
	u.OTPCreatedAt = time.Time{}
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return gorm.Expr("ST_PointFromText(?)", fmt.Sprintf("POINT(%.8f %.8f)", loc.Longitude, loc.Latitude))
}

func (loc *Location) Scan(v interface{}) error {
	hexWKBString, ok := v.(string)
	if !ok {
		return nil
	}

	decodedHexWKBString, err := ewkbhex.Decode(hexWKBString)
	if err != nil {
		log.Println(err)
	}
	arrayOfCoordinates := decodedHexWKBString.FlatCoords()
	loc.Longitude, loc.Latitude = arrayOfCoordinates[0], arrayOfCoordinates[1]
	return nil
}

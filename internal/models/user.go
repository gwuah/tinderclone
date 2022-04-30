package models

import (
	//"errors"
	"time"
	"context"
	// "database/sql/driver"
	// "encoding/hex"
	//"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	//"math"
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
	ID           string      `json:"id" gorm:"default:gen_random_uuid()"`
	CountryCode  string      `json:"country_code"`
	PhoneNumber  string      `json:"phone_number"`
	OTP          string      `json:"otp"`
	RawOTP       string      `json:"raw_otp"`
	CreatedAt    time.Time   `json:"created_at" sql:"type:timestamp without time zone" `
	OTPCreatedAt time.Time   `json:"otp_created_at" sql:"type:timestamp without time zone" `
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Bio          string      `json:"bio"`
	Location     Location    `json:"location"`
	DOB          time.Time   `json:"dob" sql:"type:timestamp without time zone"`
	Gender       string      `json:"gender"`
	Interests    string      `json:"interests"`
	Scores       Scores      `json:"scores" gorm:"-"`
}

func (u *User) Sanitize() {
	u.OTP = ""
	u.OTPCreatedAt = time.Time{}
}

func (loc Location) GormDataType() string {
	return "geometry"
  }
  
func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
	  SQL:  "ST_PointFromText(?)",
	  Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", loc.Longitude, loc.Latitude)},
	}
  }
  
  // Scan implements the sql.Scanner interface
  func (loc *Location) Scan(v interface{}) error {
	return nil
  }

// func (l *Location) String() string {
// 	return fmt.Sprintf("SRID=4326;POINT(%v %v)", l.Longitude, l.Latitude)
// }
// func (l *Location) GormDataType() string {
// 	return "GEOGRAPHY"
// }

// func (l *Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
// 	return clause.Expr{
// 		SQL:  "?",
// 		Vars: []interface{}{fmt.Sprintf("SRID=4326;POINT(%v %v)", l.Longitude, l.Latitude)},
// 	}
// }

// func (l *Location) Scan(value interface{}) error {
// 	//return nil
// 	data, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("invalid data type")
// 	}

// 	return errors.New("invalid geo data")
// }



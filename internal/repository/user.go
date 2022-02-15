package repository

import (
	"net/http"

	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db  *gorm.DB
	sms *lib.SMS
}

func NewUserRepo(db *gorm.DB, sms *lib.SMS) *UserRepo {
	return &UserRepo{db: db, sms: sms}
}

func (u *UserRepo) FindUserByPhone(phone string) (*models.User, int64, error) {
	var user models.User
	db := u.db.Where("phone_number = ?", phone).Find(&user)
	if db.Error != nil {
		return nil, db.RowsAffected, db.Error
	}
	return &user, db.RowsAffected, db.Error
}

func (u *UserRepo) CreateUser(user *models.User) error {
	return u.db.Create(&user).Error
}

func (u *UserRepo) FindUserByID(id string) (*models.User, error) {
	var user models.User
	db := u.db.Where("id = ?", id).Find(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, db.Error
}

func (u *UserRepo) SendSMS(message *lib.Message, phone_number string) (*http.Response, error) {
	message.To = phone_number
	resp, err := u.sms.SendTextMessage(*message)
	return resp, err
}

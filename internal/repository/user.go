package repository

import (
	"fmt"

	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
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
	return u.db.Omit("Location").Create(&user).Error
}

func (u *UserRepo) FindUserByID(id string) (*models.User, error) {
	var user models.User
	db := u.db.Where("id = ?", id).Find(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, db.Error
}

func (u *UserRepo) FindUserInterests(id string) ([]string, error) {
	type Interests struct {
		Interests string
	}
	var userInt Interests
	db := u.db.Where("id = ?", id).Find(&userInt)
	if db.Error != nil {
		return nil, db.Error
	}
	interests := lib.StringToSlice(userInt.Interests)

	return interests, nil
}

func (u *UserRepo) UpdateUserByID(id string, user *models.User) error {
	return u.db.Model(models.User{}).Where("id = ?", id).Updates(&user).Error
}

func (u *UserRepo) UpdateLocationByID(id string, loc models.Location) error {
	return u.db.Exec(fmt.Sprintf(`UPDATE users
	SET location = 'POINT(%.8f %.8f)'
	WHERE id = ?`,
		loc.Longitude,
		loc.Latitude),
		id,
	).Error
}

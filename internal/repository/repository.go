package repository

import (
	"github.com/gwuah/tinderclone/internal/lib"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo *UserRepo
}

func New(db *gorm.DB, sms *lib.SMS) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(db, sms),
	}
}

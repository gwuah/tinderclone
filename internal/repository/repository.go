package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo *UserRepo
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(db),
	}
}

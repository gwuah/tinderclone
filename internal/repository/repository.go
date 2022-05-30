package repository

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo *UserRepo
	RedisClient *redis.Client
}

func New(db *gorm.DB, client *redis.Client) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(db),
		RedisClient : client,
	}
}

package redistest

import (
	"os"

	"github.com/go-redis/redis"
)

func Init() (*redis.Client, error) {

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "127.0.0.1:6379"
	}
	pass := os.Getenv("REDIS_PASSWORD")
	if pass == "" {
		pass = ""
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		DB:       1,
	})

	err := rdb.Set("ping", "pong", 0).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil

}

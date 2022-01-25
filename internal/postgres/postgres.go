package postgres

import (
	"fmt"
	"os"

	"github.com/gwuah/tinderclone/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConstructDatabaseURI() string {
	dburl := os.Getenv("DATABASE_URL")
	if dburl != "" {
		return dburl
	}

	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("DB_PORT")
	SSLMODE := os.Getenv("SSLMODE")
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", USER, PASSWORD, HOST, PORT, DBNAME, SSLMODE)
}

func Init() (*gorm.DB, error) {

	databaseUrl := ConstructDatabaseURI()

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

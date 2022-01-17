package postgres

import (
	"fmt"
	"os"
	"github.com/gwuah/tinderclone/cmd/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
    USER := os.Getenv("DB_USER")
    PASSWORD := os.Getenv("DB_PASS")
    HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
    PORT := os.Getenv("DB_PORT")
	SSLMODE := os.Getenv("SSLMODE")

	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	HOST, PORT, USER, PASSWORD, DBNAME, SSLMODE)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
	}
	
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	return db
}


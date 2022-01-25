package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func New() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	pathToEnvFile := fmt.Sprintf("%s/.env", dir)

	if _, err := os.Stat(pathToEnvFile); err == nil {
		return godotenv.Load(pathToEnvFile)
	}

	return nil
}

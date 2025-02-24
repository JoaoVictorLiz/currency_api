package services

import (
	"errors"
	"os"
	"github.com/joho/godotenv"
)

func Loadenv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("It was not possible to read .ENV")
	}

	return nil
}

func GetBaseURL() string {
	return os.Getenv("BASE_URL")
}
package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitConfig() error {
	var err error
	db, err = InitDabatase()

	if err != nil {
		return fmt.Errorf("Postgres error initialize: %v", err)
	}

	return nil

}

func Getdatabase() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {

	// Initialize logger
	logger = NewLogger(prefix)
	return logger

}

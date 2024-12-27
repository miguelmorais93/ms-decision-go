package config

import (
	"database/sql"
	"fmt"
)

var (
	db     *sql.DB
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

func Getdatabase() *sql.DB {
	return db
}

func GetLogger(prefix string) *Logger {

	// Initialize logger
	logger = NewLogger(prefix)
	return logger

}

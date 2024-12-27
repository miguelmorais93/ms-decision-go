package config

import (
	"fmt"
	"ms-decision-go/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUser     = "myuser"
	DBPassword = "mypassword"
	DBName     = "mydatabase"
	DBHost     = "localhost"
	DBPort     = "5432"
)

func InitDabatase() (*gorm.DB, error) {
	logger := GetLogger("Postgres database: ")
	logger.Debugf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		DBHost, DBUser, DBPassword, DBName, DBPort)
	// String de conexão para PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		DBHost, DBUser, DBPassword, DBName, DBPort)

	// Abre a conexão com o banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Postgres error connection: %v", err)
		return nil, err
	}

	// Auto-migração para criar tabelas
	err = db.AutoMigrate(&model.DecisionUser{})
	if err != nil {
		return nil, err
	}

	logger.Infof("Postgress connected: %v", dsn)
	return db, nil
}

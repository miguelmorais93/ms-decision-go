package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DBUser     = "myuser"
	DBPassword = "mypassword"
	DBName     = "mydatabase"
	DBHost     = "localhost"
	DBPort     = "5432"
)

func InitDabatase() (*sql.DB, error) {
	logger := GetLogger("Postgres database: ")
	logger.Debugf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		DBHost, DBUser, DBPassword, DBName, DBPort)
	// String de conexão para PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		DBHost, DBUser, DBPassword, DBName, DBPort)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Errorf("Postgres error connection: %v", err)
		return nil, err
	}

	// Verifica a conexão
	err = db.Ping()
	if err != nil {
		logger.Errorf("Postgres error connection: %v", err)
		return nil, err
	}

	logger.Infof("Postgress connected: %v", dsn)
	return db, nil
}

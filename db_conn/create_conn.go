package db_conn

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// createConn createConn
func createConn() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, errors.New("DB_HOST")
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, errors.New("DB_USER")
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		return nil, errors.New("DB_PASSWORD")
	}
	dbDatabase := os.Getenv("DB_DATABASE")
	if dbDatabase == "" {
		return nil, errors.New("DB_DATABASE")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, errors.New("DB_PORT")
	}
	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbSSLMode == "" {
		return nil, errors.New("DB_SSLMODE")
	}

	dataSource := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbDatabase,
		dbSSLMode)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	return db, nil
}

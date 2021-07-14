package adapters

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

var Db *sqlx.DB

func InitDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))

	if err != nil {
		log.Printf("error while connecting: %s", err)
		return nil, errors.New("DATABASE_CONNECTING_ERROR")
	}

	if err = db.Ping(); err != nil {
		log.Printf("error while checking connection: %s", err)
		return nil, errors.New("DATABASE_CONNECTING_ERROR")
	}
	return db, nil
}

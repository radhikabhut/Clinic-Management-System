package db

import (
	"clinic-management/config"
	"log"

	"github.com/jackc/pgx"
)

var conn *pgx.Conn

func ConnectDB() error {
	cfg := config.RuntimeConfig.Postgres

	var err error
	conn, err = pgx.Connect(pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(cfg.Port),
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.DBName,
	})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return err
	}
	log.Println("Connected to the database successfully!")

	return nil
}

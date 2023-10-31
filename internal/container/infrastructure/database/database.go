package database

import (
	"fmt"
	"fww-core/internal/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var conn *sqlx.DB

func initConnection(cfg *config.DatabaseConfig) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)

	var err error
	conn, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// set connection pool
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(5)

	// ping
	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	return conn

}

func GetConnection(cfg *config.DatabaseConfig) *sqlx.DB {
	if conn == nil {
		initConnection(cfg)
	}
	return conn
}

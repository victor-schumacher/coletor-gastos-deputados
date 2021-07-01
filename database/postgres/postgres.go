package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

var (
	databaseURL = os.Getenv("DATABASE_URL")
)

const BatchSize = 500

type PgManager struct {
}

func NewPgManager() PgManager {
	return PgManager{}
}

func (p PgManager) ConnectHandle() *sql.DB {
	if databaseURL == "" {
		log.Fatal("Error: Invalid DATABASE_URL value.")
	}
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxIdleConns(32)
	db.SetMaxOpenConns(64)
	db.SetConnMaxIdleTime(time.Minute * 2)

	return db
}

func (p PgManager) TestConnection() {
	c := p.ConnectHandle()
	err := c.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

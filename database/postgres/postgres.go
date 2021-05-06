package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	driveName         = "postgres"
	dataSourcePattern = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
)

var (
	host     = os.Getenv("DATABASE_HOST")
	port     = os.Getenv("DATABASE_PORT")
	dbname   = os.Getenv("DATABASE_NAME")
	user     = os.Getenv("DATABASE_USER")
	password = os.Getenv("DATABASE_PASSWORD")
	sslMode  = os.Getenv("DATABASE_SSL_MODE")
)

type PgManager struct {
}

func NewPgManager() PgManager {
	return PgManager{}
}

func (p PgManager) ConnectHandle() *sql.DB {
	db, err := sql.Open(driveName, p.dataSource())
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

func (p PgManager) dataSource() string {
	dbPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Error: Invalid DATABASE_PORT value.")
	}

	return fmt.Sprintf(dataSourcePattern, host, dbPort, user, password, dbname, sslMode)
}

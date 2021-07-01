package database

import (
	"database/sql"
)

type DBConnection interface {
	ConnectHandle() *sql.DB
	TestConnection()
}

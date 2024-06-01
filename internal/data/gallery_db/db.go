package gallery_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	host     string
	port     int
	username string
	password string
	dbname   string
}

func NewDB(
	host string,
	port int,
	username string,
	password string,
	dbname string) *DB {
	return &DB{
		host:     host,
		port:     port,
		username: username,
		password: password,
		dbname:   dbname,
	}
}

func (d *DB) getConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.username, d.password, d.dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, err
}

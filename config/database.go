package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/url"
	"time"
)

func InitDB() (*sql.DB, error) {
	user := C.Database.DBUSER
	pass := C.Database.DBPASS
	port := C.Database.DBPORT
	host := C.Database.DBHOST
	name := C.Database.DBNAME

	val := url.Values{}
	val.Add("TimeZone", "Asia/Jakarta")

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%v sslmode=disable %s`, host, user, pass, name, port, val.Encode())

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return db, nil
}
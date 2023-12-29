package main

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:wth_PasswordAdmin1@tcp(localhost:3306)/recordings")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(50 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

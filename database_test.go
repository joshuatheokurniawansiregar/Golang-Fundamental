package main

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:wth_PasswordAdmin1@tcp(localhost:3306)/recordings")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func TestAnotherDB(t *testing.T) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordingss",
	}
	// Get a database handle.

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
}

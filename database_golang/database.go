package database_golang

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db_connection, err := sql.Open("mysql", "root:NewPassword@1@tcp(127.0.0.1:3306)/database_in_golang?parseTime=true")
	if err != nil {
		fmt.Println("Error occured: ", err.Error())
	}
	db_connection.SetMaxIdleConns(5)
	db_connection.SetMaxOpenConns(15)
	db_connection.SetConnMaxIdleTime(6 * time.Second)
	db_connection.SetConnMaxLifetime(10 * time.Minute)
	return db_connection
}

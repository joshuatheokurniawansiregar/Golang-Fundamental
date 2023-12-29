package database_golang_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func IsValidDataType(data_type interface{}, expected_datatype string) (bool, error) {
	var data_type_slice []string = make([]string, 6)
	data_type_slice = []string{"int32", "int64", "int", "string", "float32", "float64"}
	var is_true bool
	var error_message error
	var counter_not_equal int8 = 0
	for _, each := range data_type_slice {
		if expected_datatype != each {
			counter_not_equal += 1
		} else {
			if counter_not_equal == 1 {
				break
			}
		}
	}
	if counter_not_equal == 1 {
		is_true, error_message = false, errors.New("There is no data type named: "+expected_datatype)
	} else {
		is_true, error_message = true, nil
	}
	var reflect_value = reflect.ValueOf(data_type)

	switch {
	case reflect_value.Kind() == reflect.Int32 && data_type_slice[0] == expected_datatype:
		return true, nil
	case reflect_value.Kind() == reflect.Int64 && data_type_slice[1] == expected_datatype:
		return true, nil
	case reflect_value.Kind() == reflect.Int && data_type_slice[2] == expected_datatype:
		return true, nil
	case reflect_value.Kind() == reflect.String && data_type_slice[3] == expected_datatype:
		return true, nil
	case reflect_value.Kind() == reflect.Float32 && data_type_slice[4] == expected_datatype:
		return true, nil
	case reflect_value.Kind() == reflect.Float64 && data_type_slice[1] == expected_datatype:
		return true, nil
	default:
		is_true, error_message = false, errors.New(" "+expected_datatype)
	}
	return is_true, error_message
}

func TestIsValidDataType(t *testing.T) {

	is_true, err := IsValidDataType(4, "string")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(is_true)
}

func DSN(database_name string) string {
	var dsn string = "root:NewPassword@1@tcp(127.0.0.1:3306)/" + database_name
	return dsn
	// return fmt.Sprintf("%s:%s@%s(%s:%s)/%s", "root", "NewPassword@1", "tcp", "127.0.0.1", "3306", database_name)
}

func TestDatabaseConnection(t *testing.T) {
	var db_name string = "database_in_golang"
	dbconn, err := sql.Open("mysql", DSN(db_name))
	if err != nil {
		log.Printf("%s", err.Error())
	}
	ctx_timeout, cancelfunc := context.WithTimeout(context.Background(), time.Second*10)

	res, err := dbconn.ExecContext(ctx_timeout, "CREATE DATABASE IF NOT EXISTS "+db_name)
	if err != nil {
		log.Printf("Error: %s when creating connect to database", err)

	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
	}
	dbconn.SetConnMaxLifetime(3 * time.Minute)
	dbconn.SetMaxOpenConns(6)
	dbconn.SetMaxIdleConns(2)
	defer cancelfunc()
	defer dbconn.Close()
	fmt.Println(rows)
}

func GetConnectionTest() *sql.DB {
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

func TestSQLInjection(t *testing.T) {
	username := "admin'; #"
	password := "salah"

	var connection *sql.DB = GetConnectionTest()
	defer connection.Close()
	var context context.Context = context.Background()
	query := "SELECT `username`, `password` FROM `user` where `username` = '" + username + "'" + " AND `password` = '" + password + "'" + " LIMIT 1;"
	rows, queryerr := connection.QueryContext(context, query)
	if queryerr != nil {
		panic(queryerr)
	}
	fmt.Println(query)
	if rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("Aye, berhasil login")
	} else {
		fmt.Println("Gagal login")
	}

	defer rows.Close()
}

func TestSQLInjectionSafe(t *testing.T) {
	username := "admin; DROP TABLE student; #"
	password := "admin"

	var connection *sql.DB = GetConnectionTest()
	defer connection.Close()
	var context context.Context = context.Background()
	var queryInsert string = "INSERT INTO `user`(`username`, `password`) VALUES(?, ?) "
	// var querySelect string = "SELECT `username`, `password` FROM `user` where `username` = ? AND `password` = ? LIMIT 1"
	_, queryerr := connection.ExecContext(context, queryInsert, username, password)
	if queryerr != nil {
		panic(queryerr)
	}

	// if rows.Next() {
	// 	var username string
	// 	var password string
	// 	err := rows.Scan(&username, &password)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Aye, berhasil login")
	// }
	// fmt.Println("Gagal login")

	// defer rows.Close()
}

func TestAutoIncrement(t *testing.T) {
	var email string = "joshua@email.com"
	var comment string = "this is comment!"
	var connection *sql.DB = GetConnectionTest()
	defer connection.Close()
	res, err := connection.ExecContext(context.Background(), "INSERT INTO `comments`(`email`, `comment`) VALUES(?, ?)", email, comment)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

func PrepareStament(mysql_connection *sql.DB, query string) *sql.Stmt {
	defer mysql_connection.Close()
	stmt, err := mysql_connection.PrepareContext(context.Background(), query)
	if err != nil {
		panic(err)
	}
	return stmt
}

func TestPrepareStament(t *testing.T) {
	connection := GetConnectionTest()

	defer connection.Close()
	statement, err := connection.PrepareContext(context.Background(), "INSERT INTO `comments`(`email`, `comment`) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	defer statement.Close()
	for i := 1; i <= 10; i++ {
		var inttostring string = strconv.Itoa(i)
		email := "email" + inttostring + "@gmail.com"
		comment := "comment " + inttostring
		result, err := statement.ExecContext(context.Background(), email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil: ", id)
	}
}

func TestTransaction(t *testing.T) {
	mysql_connection := GetConnectionTest()
	defer mysql_connection.Close()

	context := context.Background()
	tx, err := mysql_connection.BeginTx(context, nil)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	var email string = "this is string email"
	var comments string = "this is string comment"
	var query string = "INSERT INTO `comments`(`email`, `comment`) VALUES(?, ?)"
	stmt, err := tx.PrepareContext(context, query)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	is_string_email, _ := IsValidDataType(email, "string")
	is_string_comment, _ := IsValidDataType(comments, "string")
	if is_string_email && is_string_comment {
		result, err := stmt.ExecContext(context, email, comments)
		if err != nil {
			fmt.Errorf("%s", err)
		}
		id, err := result.LastInsertId()
		if err == nil {
			fmt.Errorf("%s", err)
		}

		if err = tx.Commit(); err != nil {
			fmt.Errorf("%s", err.Error())
		}
		fmt.Println(id)
	} else {
		if err = tx.Rollback(); err != nil {
			fmt.Errorf("%s", err.Error())
		}
		fmt.Println(err)
	}
}

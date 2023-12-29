package main

import (
	// "Basics/Fundamentals/algorithms_and_data_structures"
	// "Basics/Fundamentals/algorithms_and_data_structures/data_structures"

	"Fundamentals/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

func createfile() {
	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(file, "data")
	fmt.Println("file created")
}

func CreateFile() {
	// create directory on os and file on the directory

	err := os.Mkdir("admin_dir", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("test.txt", []byte("hello, world"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	//get last modified time
	// filename := "aye.txt"
	// fileInfo, err := os.Stat("dir/" + filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// mTime := fileInfo.ModTime()
	// fmt.Println(mTime)
	// get file size
	// var fsize int64 = fileInfo.Size()
	// fmt.Printf("the size is %d bytes\n", fsize)

	// read file
	// data, err := os.ReadFile("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// str := string(data)
	// fmt.Println(str)

}

type NotFoundError struct {
	message string
}

func (notfounderror *NotFoundError) Error() string {
	return notfounderror.message
}

func TestError() error {
	var name string = "joshua"
	if name != "Joshua" {
		return &NotFoundError{"user not found"}
	}
	return nil
}

type Customer struct {
	id                   int
	name                 string
	email                sql.NullString
	balance              int64
	rating               float64
	created_at, birth_at sql.NullTime
	married              bool
}

var customerarray []Customer = []Customer{}

//	func InsertInto(table_name string, column_params ...any) {
//		var map_column_params map[int]any
//		var query string = ""
//		var length int = len(column_params)
//		for {
//			i := 1
//			map_column_params[i] = column_params
//			if i == length {
//				break
//			}
//		}
//		for _, data := range map_column_params {
//			query = "INSERT INTO " + "`" + table_name + "`" + "VALUES(" + column_params[i] + ")"
//		}
//	}
func IsValidDataType(data_type interface{}, expected_datatype string) (bool, error) {
	var data_type_slice []string = make([]string, 6)
	data_type_slice = []string{"int32", "int64", "int", "string", "float32", "float64", "bool"}
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
	case reflect_value.Kind() == reflect.Float64 && data_type_slice[5] == expected_datatype:
		return true, nil
	case reflect_value.Kind() == reflect.Float64 && data_type_slice[6] == expected_datatype:
		return true, nil
	default:
		is_true, error_message = false, errors.New("Expected data type `"+expected_datatype+"` is not same with the first parameter")
	}
	return is_true, error_message
}

func main() {
	// is_true, err := IsValidDataType(true, "string")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(is_true)
	ctx := context.Background()
	repository_comment := repository.ConstructorCommentInterfaceImpl(ctx)

	all_comments, err := repository_comment.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(all_comments)

	// connection := database_golang.GetConnection()
	// defer connection.Close()
	// var context context.Context = context.Background()
	// rows, queryerr := connection.QueryContext(context, "SELECT `id`, `name`, `email`, `balance`, `rating`, `created_at`, `birth_at`, `married` FROM `customer`")
	// if queryerr != nil {
	// 	panic(queryerr)
	// }

	// var id int
	// var name string
	// var email sql.NullString
	// var balance int64
	// var rating float64
	// var created_at sql.NullTime
	// var birth_at sql.NullTime
	// var married bool
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_at, &married)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	customerarray = append(customerarray, Customer{id: id, name: name, email: email, balance: balance, rating: rating, created_at: created_at, birth_at: birth_at, married: married})
	// }
	// fmt.Println(id)
	// fmt.Println(customerarray)
	// defer rows.Close()

	// ctx, cancelcontext := context.WithTimeout(context.Background(), 3*time.Second)
	// var db_connection *sql.DB = database_golang.GetConnection()
	// defer db_connection.Close()
	// _, err := db_connection.ExecContext(ctx, "INSERT INTO `customer`(`id`, `name`) VALUES(1, 'Joshua', );")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer cancelcontext()
	// fmt.Println("Success to insert customer")

	// err_ := TestError()
	// if err_ != nil {
	// 	fmt.Println("hai")
	// 	fmt.Println(err_)
	// 	if test_custom_error, ok := err_.(*NotFoundError); ok {
	// 		fmt.Printf("not found error: %v", test_custom_error.Error())
	// 	}
	// }
	// standard_library.RunningCSVReading()

	// standard_library.CreateNewFile("tmp/ayo1.txt", "\nHI! WORLD")

	// createfile()tmp\defer.txt
	// CallsInterfaceKosong()
	// CallChannelFunction()
	// var tnode *data_structures.Treenode = nil
	// tnode = data_structures.InsertTree(tnode, 50)
	// tnode = data_structures.InsertTree(tnode, 30)
	// tnode = data_structures.InsertTree(tnode, 20)
	// tnode = a.InsertTree(tnode, 40)
	// tnode = a.InsertTree(tnode, 70)
	// tnode = a.InsertTree(tnode, 60)
	// tnode = a.InsertTree(tnode, 80)
	// tnode = a.InsertTree(tnode, 100)
	// tnode = a.InsertTree(tnode, 90)
	// tnode = a.InsertTree(tnode, 35)
	// tnode = a.InsertTree(tnode, 32)
	// tnode.InorderTree(tnode)

	// var BinarySearchTree a.BinarySearchTree
	// BinarySearchTree.Insert(2)
	// BinarySearchTree.Insert(3)
	// newBSTRoot := a.NewRoot(&BinarySearchTree)
	// BinarySearchTree.InorderTree(newBSTRoot)

	// pkg.CallsRecover()
	// pkg.CallsPointers()
	// pkg.CallsRecover()
	// pkg.CallsMutVars()
	// pkg.CallsImutVars()
	// pkg.CallTypeAssertion()
	// pkg.CallsReflectIsValid()

	// n := 2
	// var array = []int{3, 2, 4, 5}
	// fmt.Println(algorithms_and_data_structures.TwoSum(array, 7))
	// fmt.Println(algorithms_and_data_structures.TwoSumWithTwoPointer(array, 7))
	// fmt.Println(algorithms_and_data_structures.Anagram("rat", "car"))

	// var testreflect = reflect.ValueOf(CallRuneVariableFunction())
	// fmt.Println("test: ", testreflect.Kind())
	// CallsReflect()

}

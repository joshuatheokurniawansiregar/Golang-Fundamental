package pkg

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true"`
}

type User struct {
	Name string `required:"true"`
	Age  int    `required:"true"`
}

type testfunc func(string) bool

func TestReflect() {
	var daa int = 10
	var reflectvalue = reflect.ValueOf(daa)
	var reflectKind reflect.Kind = reflectvalue.Kind()
	switch reflectKind {
	case reflectKind:

	}
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	var isValid bool = true
	var valueArray = make(map[int]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			valueArray[i] = value

		}
	}
	for _, each := range valueArray {
		switch type_assert := each.(type) {
		case string:
			if type_assert == "" {
				isValid = false
			}
		case int:
			if type_assert == 0 {
				isValid = false
			}
		default:
			isValid = true
		}
	}
	return isValid
}

func CallsReflectIsValid() {
	var usersample User = User{}
	usersample.Name = "sdf"
	usersample.Age = 9
	fmt.Println(IsValid(usersample))
}

func CallsReflect() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe data variable ", reflectValue)

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable = ", reflectValue.Int())
	}
}

package pkg

import "fmt"

//User-defined types
//A type is used for making method in go lang
type ints int
type stringss string
type object struct {
	attribute_string1, attribute_string2     string
	attribute_int8_1, attribute_int8_2       int8
	attribute_float32_1, attribute_float32_2 float32
	attribute_boolean                        bool
}
type arrays []int8
type string_array []string

var string_helloworld = "hello world"
var object_array = []object{
	{attribute_string1: "hello", attribute_string2: "world", attribute_int8_1: 12},
	{attribute_boolean: false},
}

func CallRuneVariableFunction() int32 {
	var identity = func(v rune) rune { return v }
	var i int32 = 10
	return identity(i)
}

func VariablesFunc() {
	// Int: int, int8, int16, int32, int64
	// Float: float32, float64
	// boolean: bool
	// String: string
	// No double
	// var a, c int16 = 222, 12
	var str string = "+-*/"
	// var dbl float32 = 1.0 + 1.3
	f, d := 1, 2
	// fmt.Println("result: ", a+c)
	fmt.Println(str[0:3])
	fmt.Println(str[1])
	fmt.Println(str[2])
	fmt.Println(str[3])
	fmt.Println(f + d)
	// fmt.Println(dbl)
}

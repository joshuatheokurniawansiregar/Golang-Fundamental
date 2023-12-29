package pkg

import "fmt"

type AyePerson struct {
	name string
	age  int
}

func CallsInterfaceKosong() {
	//casting interface{}
	fmt.Println("casting variable interface to another data type")
	var variable interface{}
	variable = 1
	fmt.Println(variable)
	variable = true
	fmt.Println(variable)
	variable = "satu"
	fmt.Println(variable)

	//casting variable interface to object interface
	fmt.Println("\n casting variable interface to object interface")
	var secret interface{}
	secret = &AyePerson{name: "Joshua Theo Kurniawan", age: 23}
	var name string = secret.(*AyePerson).name
	fmt.Println(name)

	//array interface
	fmt.Println("\n array interface")
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	for _, value := range fruits {
		fmt.Println(value)
	}
}

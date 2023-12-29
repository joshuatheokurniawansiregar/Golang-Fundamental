package pkg

import "fmt"

func CallTypeAssertion() {
	var hello interface{} = "hello"
	result, boo := hello.(string)
	fmt.Println(result, boo)
}

package pkg

import "fmt"

type Address struct {
	City, Province, Country string
}

func CallsPointers() {
	// * symbol is used to declare pointer and dereference
	// & is used to points to the address of the stored value
	var nullpointerAddress Address = Address{}
	var pointerAddress *Address = &nullpointerAddress
	*pointerAddress = Address{"Malang", "Jawa Timur", "Indonesia"}
	var pointerAddress2 *Address = &nullpointerAddress
	var RegulerAddress Address = nullpointerAddress
	fmt.Println(pointerAddress2)
	fmt.Println(RegulerAddress)
	i := 1
	ai := 4
	var a *int = &i
	var pointer *int = a
	a = &ai
	fmt.Println(" ", &i)       // showing memory address
	fmt.Println(" ", *a)       // showing the value of the dereference of the pointer: a
	fmt.Println(" ", &a)       // showing memory address of the the pointer a
	fmt.Println(" ", *pointer) // showing value of pointer :pointer
}

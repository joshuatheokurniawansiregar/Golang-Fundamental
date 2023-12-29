package pkg

import "fmt"

func CallsMutVars() {
	//Mutable data type in Golang:
	//slice
	//array
	//map
	//channels

	var slice []int = []int{1, 2, 3}
	fmt.Printf("S[1] -> %p\n", &slice[1])
	slice1 := slice
	slice1[1] = 4
	fmt.Printf("S1[1] -> %p\n", &slice1[1])
	fmt.Println("s = ", slice)
	fmt.Println("s = ", slice1)
}

type imutable struct {
	name string
}

func CallsImutVars() {
	//Mutable data type in Golang:
	//boolean
	//int
	//string
	//interface
	//pointer

	var boo bool = true
	b := boo
	b = false
	fmt.Printf("memory address of boo: %p\n", &boo)
	fmt.Printf("memory address of b: %p\n", &b)
	fmt.Println("boo = ", boo)
	fmt.Println("b = ", b)

	var ob imutable = imutable{"joshua"}
	var ob1 imutable = ob
	ob1.name = "joshua theo"
	fmt.Printf("ob's memory address : %p\n", &ob)
	fmt.Printf("ob's memory address : %p\n", &ob1)
	fmt.Println(ob)
	fmt.Println(ob1)
}

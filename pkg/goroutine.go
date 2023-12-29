package pkg

import (
	"fmt"
	"runtime"
	"time"
)

func workers(sycn chan bool) {
	fmt.Println("working")
	time.Sleep(2 * time.Second)
	fmt.Println("next")
	sycn <- true
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func CallsGoRoutine() {
	var isdone chan bool = make(chan bool)
	workers(isdone)
	fmt.Println("isdone: ", <-isdone)
	f("direct")
	go f("go routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func CallsGoRoutine1() {
	runtime.GOMAXPROCS(2)
	go print(5, "hello world")
	print(5, "hello dunia")
	var input string
	fmt.Scanln(&input)
}

package pkg

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func CallGoRoutineCurrency() {
	go say("world")
	say("hello")
}

func sum(a []int, c chan int) {
	total := 0
	for r := range a {
		total = total + r
	}
	c <- total
}

func CallChannelCorruncy() {
	a := []int{7, 2, 8, -9}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(c, y, x+y)
}

func FibonacciWithChan(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		y, x = x+y, y
	}
	close(c)
}

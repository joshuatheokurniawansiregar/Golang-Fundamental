package pkg

import (
	"fmt"
	"runtime"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func printMessage(message chan string) {
	fmt.Println(<-message)
}

func CallChannelAsParameterDataType() {
	var arrstring []string = []string{"Joshua", "Gege Akutami", "Me and Myselves"}
	message := make(chan string)
	for _, r := range arrstring {
		go func(who string) {
			var smessagedata string = "hello " + who
			message <- smessagedata
		}(r)
	}
	lengthofarrstring := len(arrstring)
	for i := 0; i < lengthofarrstring; i++ {
		printMessage(message)
	}
}

func CallChannelFunction() {
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println(<-done)
	runtime.GOMAXPROCS(2)
	var chn chan string = make(chan string)
	var sayHello = func(who string) {
		var zer_valstr string = "Hello " + who
		chn <- zer_valstr
	}
	go sayHello("Joshua Theo")
	go sayHello("Gege Akutami")
	go sayHello("Me and myselves")
	var message1 string = <-chn
	fmt.Println(message1)
	var message2 string = <-chn
	fmt.Println(message2)
	var message3 string = <-chn
	fmt.Println(message3)
}

func CallBufferedChannel() {
	// runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}

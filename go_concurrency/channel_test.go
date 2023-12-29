package go_concurrency

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	chann := make(chan string)
	defer close(chann)
	go func() {
		time.Sleep(2 * time.Second)
		chann <- "Hello World"
		fmt.Println("Data channel has been sent out")
	}()
	fmt.Println(<-chann)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Joshua Theo Kurniawan"
}

func TestChannelAsParameter(t *testing.T) {
	var channel chan string = make(chan string)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func ChanIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Sent Channel"

}

func ChanOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChanInOut(t *testing.T) {
	var channel chan string = make(chan string)
	defer close(channel)
	go ChanIn(channel)
	go ChanOut(channel)
	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	var channel chan string = make(chan string, 4)
	defer close(channel)
	fmt.Println(runtime.NumGoroutine())
	go func() {
		channel <- "Joshua"
		channel <- "Theo"
		channel <- "Kurniawan"
		channel <- "Siregar"
		fmt.Println("goroutine 1")
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println("goroutine 2")
	}()

	go func() {
		for i := range channel {
			fmt.Println(i)
		}
		fmt.Println("goroutine 3")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
	fmt.Println(runtime.NumGoroutine())
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		i := 0
		for {
			channel <- "perulangan ke" + strconv.Itoa(i+1)
			if i == 9 {
				break
			}
			i++
		}
		defer close(channel)
	}()
	for data := range channel {
		fmt.Println("Menerima data " + data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter_loop := 0
	for {
		select {
		case data := <-channel1:
			counter_loop++
			fmt.Println("Data dari channel 1: ", data)
		case data := <-channel2:
			counter_loop++
			fmt.Println("Data dari channel 2: ", data)
		}
		if counter_loop == 2 {
			break
		}
	}
}

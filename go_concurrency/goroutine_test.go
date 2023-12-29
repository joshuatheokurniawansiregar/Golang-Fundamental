package go_concurrency

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}
func TestGoRoutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)

	}
	time.Sleep(5 * time.Second)
}

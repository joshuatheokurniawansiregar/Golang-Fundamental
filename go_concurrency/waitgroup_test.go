package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncrounous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	var group sync.WaitGroup = sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsyncrounous(&group)
	}
	group.Wait()
	fmt.Println("Selesai")
}

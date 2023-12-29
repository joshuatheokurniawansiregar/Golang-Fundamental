package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnlyOnce(t *testing.T) {
	var once *sync.Once = &sync.Once{}
	var group *sync.WaitGroup = &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(OnlyOnce)
		}()
	}
	group.Wait()
	fmt.Println("counter: ", counter)
}

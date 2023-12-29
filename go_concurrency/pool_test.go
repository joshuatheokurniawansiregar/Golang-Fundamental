package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return "Hello World"
		},
	}
	var wgroup *sync.WaitGroup = &sync.WaitGroup{}
	pool.Put("Joshua")
	pool.Put("23")
	pool.Put("poor")
	for i := 0; i < 10; i++ {
		wgroup.Add(1)
		go func() {
			defer wgroup.Done()
			var data interface{} = pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	wgroup.Wait()
}

package go_context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	go_context_background := context.Background()
	go_context_todo := context.TODO()
	fmt.Println(go_context_background)
	fmt.Println(go_context_todo)
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		counter := 1
		defer close(destination)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func CreateCounterWithWaitGroup(ctx context.Context, wgroup *sync.WaitGroup) chan int {
	destination := make(chan int)
	wgroup.Add(1)

	go func() {
		counter := 1
		defer close(destination)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	defer wgroup.Done()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Gouroutine ", runtime.NumGoroutine())
	var parentcontext = context.Background()
	ctx, cancels := context.WithCancel(parentcontext)
	var destination chan int = CreateCounter(ctx)
	fmt.Println("Gouroutine ", runtime.NumGoroutine())
	for i := range destination {
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
	cancels()
	fmt.Println(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("Gouroutine ", runtime.NumGoroutine())
}

func TestContextWithCancelWithWaitGroup(t *testing.T) {
	fmt.Println("Gouroutine ", runtime.NumGoroutine())
	var parentcontext = context.Background()
	var wGroup = sync.WaitGroup{}
	ctx, cancels := context.WithCancel(parentcontext)
	var destination chan int = CreateCounterWithWaitGroup(ctx, &wGroup)
	fmt.Println("Gouroutine ", runtime.NumGoroutine())
	for i := range destination {
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
	cancels()
	wGroup.Wait()
	fmt.Println("Gouroutine ", runtime.NumGoroutine())
}

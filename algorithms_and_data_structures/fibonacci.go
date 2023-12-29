package algorithms_and_data_structures

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IterativeFibonacciUsingClosure(n int) func() int {
	x, y := 1, 1
	return func() int {
		r := x
		x, y = y, x+y

		return r
	}
}

func IterativeFibonacciUsingChannel(n int, ch chan int) {
	n1, n2 := 1, 1
	for i := 0; i < n; i++ {
		ch <- n1
		n1, n2 = n2, n1+n2
	}
	close(ch)
}

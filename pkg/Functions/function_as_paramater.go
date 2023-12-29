package Functions

import "fmt"

func CallParameterAsFunction1() {
	prints(cubic, 2, 3, 3)
}

func prints(fun func(int, int, int) int, a, b, c int) {
	fmt.Println(fun(a, b, c))
	fmt.Println(c)
}

func cubic(a, b, c int) int {
	return a * b * c
}

func CallParameterAsFunction2() {
	var result string
	result = IntToString(123)
	fmt.Println(result)

	result = quote(IntToString)
	result = quote(func(a int) string { return fmt.Sprintf("%q", a) })
	fmt.Println(result)
}

type convertfn func(int) string

func IntToString(x int) string {
	return fmt.Sprintf("%v", x)
}

func quote(fn convertfn) string {
	value := 1
	return fmt.Sprintf("%q", fn(value))
}

func RunQuote() {
	quote := quote(func(x int) string { return IntToString(x) })
	fmt.Println(quote)
}

func Add(value int) int {
	return value + value
}

func CallParameterAsFunction3() {
	arr := []int{1, 2, 3, 4}
	var test []int = mapper(Add, arr)
	fmt.Println(test)
}

func mapper(fn func(value int) int, arrlist []int) []int {
	var list []int = make([]int, len(arrlist), len(arrlist))
	for i, value := range arrlist {
		list[i] = fn(value)
	}
	return list
}

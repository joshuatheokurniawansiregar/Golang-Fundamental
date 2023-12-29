package Functions

import "fmt"

//Return map[string]int
func MapFunction() map[string]int {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3
	m["k4"] = 4
	return m
}

//return array integer
func IntArrayFunction() []int {
	nums := []int{1, 2, 3, 4}
	return nums
}

// multiple return values
func MultipleVals() (int, string) {
	return 3, "tiga"
}

// variadic function
func VariadicSumNumbers(nums ...int) {
	fmt.Print("\n")
	fmt.Println(nums)
	total := 0
	// for _, num := range nums {
	// 	total += num
	// }
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	fmt.Println(total)
}

//Closure
//Go supports anonymous function, which can
//form closures.
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//function: pass by value.
//ZeroVal has integer parameter, so arguments
//will be passed to the arguments by value.
func ZeroVal(ival int) {
	ival = 0
}

//function: pass by reference.
/* ZeroPtr in contrast has an *int parameter(pointer parameter), meaning that
it takes an int pointer. The *intptr parameter in the function body then
dereferences the pointer from its memory address to current value
at that address. Assigning a value to a dereferenced pointer changes
the value at the referenced address.
*/
func ZeroPtr(intptr *int) {
	*intptr = 0
}

func CallsPointers() {
	i := 1
	fmt.Println("Before ZeroVal be called")
	fmt.Println("i = ", i)
	ZeroVal(i)
	fmt.Println("After ZeroVal be called")
	fmt.Println("i = ", i)
	fmt.Println("Before ZeroPtr be called")
	fmt.Println("i = ", i)
	ZeroPtr(&i)
	fmt.Println("After ZeroPtr be called")
	fmt.Println("i = ", i)

	VariadicSumNumbers(1, 2, 3, 4)
}

func HelloWeeb() string {
	return "Hello Weeb"
}

func CallsFunctionValue() {
	var valueFunc func() string = HelloWeeb
	var result string = valueFunc()
	fmt.Println(result)
}

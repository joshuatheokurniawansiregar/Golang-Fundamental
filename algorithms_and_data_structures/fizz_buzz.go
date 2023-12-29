package algorithms_and_data_structures

import "fmt"

func FizzBuzzzzz(numbers int) []string {
	var string_arr = make([]string, numbers)
	for i := 0; i <= numbers; i++ {
		if i%5 != 0 && i%3 != 0 {
			string_arr = append(string_arr, "fizz_buzz")
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
			string_arr = append(string_arr, "fizz")
		} else if i%3 == 0 {
			fmt.Println("buzz")
			string_arr = append(string_arr, "buzz")
		}
		// temp++
	}
	return string_arr
}

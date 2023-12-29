package pkg

import (
	"Fundamentals/pkg/Functions"
	"fmt"
)

func Range() {
	sum := 0
	var arr []int = Functions.IntArrayFunction()
	for _, num := range arr {
		sum += num
	}
	fmt.Println("sum = ", sum)
	for i, num := range arr {
		if num == 4 {
			fmt.Println("index: ", i)
		}
	}
	maps := Functions.MapFunction()
	for key, values := range maps {
		fmt.Printf("%s= %d, ", key, values)
	}

}

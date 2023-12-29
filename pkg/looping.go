package pkg

import "fmt"

func Looping() {
	// while looping
	i := 0
	for i < 3 {
		i = i + 1
		fmt.Println(i)
	}

	// General for-loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//for-each range loop
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " = ", "even")
		} else if i%2 != 0 {
			fmt.Println(i, "= ", "odd")
		}
	}
}

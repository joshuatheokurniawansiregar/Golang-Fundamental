package pkg

import "fmt"

func Arrays() {
	f := []int{1, 2, 3}

	for j := 0; j < len(f); j++ {
		fmt.Println(f[j])
	}
}

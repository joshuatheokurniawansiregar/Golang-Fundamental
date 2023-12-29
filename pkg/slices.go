package pkg

import "fmt"

func Slices() {
	var s []string
	fmt.Println("Uninit. s is null ", s, " ", s == nil, len(s) == 0)

	s = make([]string, 3, 5)
	fmt.Printf("length: %v capacity: %v", len(s), cap(s))
	// cpy:=
}

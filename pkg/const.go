package pkg

import (
	"fmt"
	"math"
)

const s string = "constant string"

func Const() {
	fmt.Println(s)
	const n float64 = 500000000
	const d float64 = 3e20 / n
	fmt.Println(d) //typecasting
	fmt.Println(math.Sin(n))
}

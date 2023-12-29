package pkg

import (
	"fmt"
	"time"
)

func Switch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday")
	}
	tme := time.Now().UnixMilli()

	fmt.Println(tme)
	// switch time.Now().Weekday() {
	// case tme.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		case int16:
			fmt.Println("I'm a int16")
		default:
			fmt.Printf("The type is %T\n", t)
		}
	}
	whatAmI(6e+11)
	whatAmI(128)
}

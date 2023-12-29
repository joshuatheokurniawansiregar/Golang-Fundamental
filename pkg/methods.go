package pkg

import "fmt"

type Rect struct {
	width, height int8
}

// This are method has a receiver type of *rect or
// pointer receiver type
func (r *Rect) Area() int8 {
	return r.width * r.height
}

// Here's an example of a value receiver
func (r Rect) Perim() int8 {
	return 2*r.width + 2*r.height
}

func CallsMethod() {
	r := Rect{width: 10, height: 5}
	fmt.Println(r.Area())
	fmt.Println(r.Perim())
}

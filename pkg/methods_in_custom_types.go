package pkg

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	width, height, depth float64
	color                Color
}

type Color byte
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
func (b *Box) SetColor(c Color) {
	b.color = c
}

func (boxlist BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE) //default
	for _, b := range boxlist {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (boxlist BoxList) PaintItBlack() {
	for i := range *&boxlist {
		boxlist[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "Yellow"}
	return strings[c]
}

func Running() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 4, 1, WHITE},
		Box{10, 10, 1, YELLOW},
	}

	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())

}

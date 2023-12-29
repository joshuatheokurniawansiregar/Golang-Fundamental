package pkg

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	r := fmt.Sprintf("base with num=%v", b.num)
	return r
}

type container struct {
	base
	str string
}

func CallsStructEmbedding() {
	var co = &container{}
	var bases *base = &base{}
	bases.num = 1
	co.base = *bases
	co.str = "some name"
	fmt.Println(co)

}

func CallsStructEmbedding2() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str:%v}\n", co.num, co.str)
	fmt.Println("Alson num= ", co.base.num)
	fmt.Println("describer", co.describe())
	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe())
}

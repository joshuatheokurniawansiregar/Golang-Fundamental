package pkg

import "fmt"

func Map() {
	var mapname map[string]int = make(map[string]int)
	var maptesting = map[string]interface{}{
		"name": "Joshua Theo Kurniawan",
	}
	fmt.Println(maptesting["name"])
	mapname["k1"] = 7
	mapname["k2"] = 13

	v1 := mapname["k1"]
	fmt.Println(v1)

	fmt.Println("Lenght ", len(mapname))
	delete(mapname, "k2")
	fmt.Println("map", mapname)

	_, prs := mapname["k1"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	n["new_key"] = 3
	fmt.Println("map foo bar", n)
}

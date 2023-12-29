package pkg

import (
	"container/list"
	"fmt"
)

func Sf() {

	var lists *list.List = list.New()
	var lastelement *list.Element

	lists.PushBack(1)
	lists.PushBack(2)
	for e := lists.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for lastelement = lists.Front(); lastelement.Value != 2; lastelement = lastelement.Next() {
	}
	lastelement = lists.InsertAfter(3, lastelement)
	lists.MoveBefore(lastelement, lists.Front())
	for e := lists.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("\n", lists.Back().Value)
}

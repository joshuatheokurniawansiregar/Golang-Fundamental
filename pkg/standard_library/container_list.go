package standard_library

import (
	"container/list"
	"fmt"
)

func ContainerList() {
	var container_list *list.List = list.New()
	container_list.PushBack(23)
	container_list.PushBack(24)
	container_list.InsertAfter(20, container_list.Back().Prev())
	for e := container_list.Front(); e != nil; e = e.Next() {
		fmt.Println(" ", e.Value)
	}
}

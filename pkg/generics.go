package pkg

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	var r = make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

//List is a singly-linked list
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(V T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: V}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: V}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func CallsGenerics() {
	var m = map[int]string{1: "2", 5: "1", 2: "2", 4: "8"}
	fmt.Println("keys: ", MapKeys(m))
	_ = MapKeys[int, string](m)

	lst := &List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(10)
	lst.Push(23)
	fmt.Println("list", lst.GetAll())
}

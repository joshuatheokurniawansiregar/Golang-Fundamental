package data_structures

import "log"

type Queue struct {
	elements []any
	len      int
}

func NewQueue(length int) *Queue {
	return &Queue{elements: nil, len: length}
}

func (q *Queue) GetLength() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return q.GetLength() == 0
}

func (q *Queue) Enque(element any) {
	if q.GetLength() == q.len {
		log.Fatal("Underflow")
		return
	}
	q.elements = append(q.elements, element)
}

// func (q *Queue) Deque() any {
// 	if q.IsEmpty() {
// 		log.Fatal("Underflow")
// 		return nil
// 	}
// 	var element = q.elements[0]
// 	if q.GetLength() == 1 {
// 		q.elements = nil
// 		return element
// 	}
// 	q.elements = q.elements[1:]
// 	return element
// }

func (q *Queue) Queue() []any {
	return q.elements
}

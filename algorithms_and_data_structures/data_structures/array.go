package data_structures

import (
	"errors"
)

type ArrayGenerics struct {
	data   []int
	length uint
}

func NewArray(capacity uint) *ArrayGenerics {
	if capacity == 0 {
		return nil
	}
	return &ArrayGenerics{
		data:   make([]int, capacity),
		length: 0,
	}
}

func (a *ArrayGenerics) Len() uint {
	return a.length
}

func (a *ArrayGenerics) isIndexOutOfBounds(index uint) bool {
	return index >= a.length
}

func (a *ArrayGenerics) InsertInteger(index uint, value int) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("full array")
	}
	if index != a.length && a.isIndexOutOfBounds(index) {
		return errors.New("out of index range")
	}
	// for i := a.length; i > index; i-- {
	// 	a.data[i] = a.data[i-1]
	// }
	a.data[index] = value
	a.length++
	return nil
}

func (a *ArrayGenerics) InsertVariadicInteger(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		a.data[i] = numbers[i]
		a.length++
	}
}

func (array *ArrayGenerics) Print() []any {
	var elems []any
	for i := 0; i < int(array.Len()); i++ {
		elems = append(elems, array.data[i])
	}
	return elems
}

func (a *ArrayGenerics) Delete(index uint) (int, error) {
	if a.isIndexOutOfBounds(index) {
		return 0, errors.New("out of index")
	}
	value := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return value, nil
}

func (a *ArrayGenerics) Swap(v1, v2 *int) {
	temp := *v1
	*v1 = *v2
	*v2 = temp
}

func (a *ArrayGenerics) BubbleSortArrayAsc() error {
	//3129457
	is_swaped := false
	for i := 0; i < int(a.length)-1; i++ {
		for j := 0; j < int(a.length)-i-1; j++ {
			// if a.length <= 1 {
			// 	return errors.New("out of range")
			// }
			if a.data[j] > a.data[j+1] {
				a.Swap(&a.data[j], &a.data[j+1])
				is_swaped = true
			}
		}
		if is_swaped == false {
			break
		}
	}
	return nil
}

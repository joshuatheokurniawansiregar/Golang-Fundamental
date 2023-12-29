package data_structures

import (
	"testing"
)

func TestInsertInteger(t *testing.T) {
	// var IntegerArray = ArrayGenerics[int]{}
	capacity := 10
	objIntegerArray := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := objIntegerArray.InsertInteger(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	objIntegerArray.Print()
}
func TestDelete(t *testing.T) {
	capacity := 10
	objIntegerArray := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := objIntegerArray.InsertInteger(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}

	for i := 1; i <= 3; i++ {
		_, err := objIntegerArray.Delete(uint(i))
		if nil != err {
			t.Fatal(err.Error())
		}
		objIntegerArray.Print()
	}
}

func TestBubbleSortArray(t *testing.T) {
	array := NewArray(1)
	err := array.InsertInteger(uint(0), 1)
	if nil != err {
		t.Fatal(err.Error())
	}
	err_sort := array.BubbleSortArrayAsc()
	if nil != err_sort {
		t.Fatal(err.Error())
	}
}

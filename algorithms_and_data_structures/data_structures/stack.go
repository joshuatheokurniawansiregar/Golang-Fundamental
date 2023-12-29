package data_structures

import (
	"errors"
	"fmt"
)

// Stack: first in last out (FILO) & last in last in first out (LIFO)
const max int = 10

type stack_any struct {
	data   [max]int
	top    int
	length int
}

func InitStackAny() *stack_any {
	return &stack_any{top: -1, length: 0}
}

func (s *stack_any) Push(v int) error {
	if s.top >= (max - 1) {
		return errors.New("stack overflow exception")
	} else {
		s.top++
		s.length++
		s.data[s.top] = v
		return nil
	}
}

func (s *stack_any) Pop() (int, error) {
	if s.top < 0 {
		return -1, errors.New("stack overflow exception")
	} else {
		s.top--
		s.length--
		var item int = s.data[s.top+1]
		return item, nil
	}
}

func (s *stack_any) Top() (int, error) {
	if s.top < 0 {
		return -1, errors.New("stack overflow exception")
	} else {
		var item int = s.data[s.top]
		return item, nil
	}
}

func (s *stack_any) YieldLength() int {
	return s.length
}

func (s *stack_any) YieldtCurrentTop() int {
	return s.top
}

func (s *stack_any) Aye() []int {
	return s.data[:s.YieldLength()]
}

func (s *stack_any) Empty() bool {
	return (s.YieldLength()) == 0
}

type Stack []string

func NewStack() *Stack {
	test := make(Stack, 0)
	return &test
}

func (st *Stack) Push(value string) {
	*st = append(*st, value)
}

func (st *Stack) Length() int {
	return len(*st)
}

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Pop() (string, bool) {
	if st.IsEmpty() {
		return "", false
	} else {
		var index int = len(*st) - 1
		var element string = (*st)[index]
		*st = (*st)[:index]
		return element, true
	}
}
func (st *Stack) Top() (string, bool) {
	if st.IsEmpty() {
		return "", false
	} else {
		var index int = len(*st) - 1
		var element string = (*st)[index]
		// *st = (*st)[:index]
		return element, true
	}

}

func (st *Stack) Print() {
	for !st.IsEmpty() {
		x, z := st.Pop()
		if z {
			fmt.Println(x)
		}
	}
}

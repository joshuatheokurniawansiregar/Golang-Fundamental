package data_structures

import (
	"errors"
	"strings"
)

type Stack_Implementation[T any] []T
type StringLib string

func (string_lib *StringLib) Substr(start, end int) string {
	*string_lib = (*string_lib)[start:end]
	var result string
	var per_byte_word []byte = make([]byte, len(*string_lib))
	n := cap(per_byte_word)
	var j int = 0
	for i := 0; i < n; i++ {
		var string_to_byte byte = (*string_lib)[i]
		per_byte_word[j] = string_to_byte
		j++
	}
	result = string(per_byte_word)
	return strings.TrimSpace(result)
}

func (stack *Stack_Implementation[any]) InsertIn(value any) {
	*stack = append(*stack, value)
}

func (stack *Stack_Implementation[any]) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack_Implementation[T]) Pop() (error, T) {
	if stack.IsEmpty() {
		var index int = len(*stack) - 1
		var top_element T = (*stack)[index]
		return errors.New("NullPointerException"), top_element
	} else {
		var index int = len(*stack) - 1
		var poped_element T = (*stack)[index]
		*stack = (*stack)[:index] //it slices the last element
		return nil, poped_element
	}

}

func (stack *Stack_Implementation[T]) Top() (error, T) {
	if stack.IsEmpty() {
		var index int = len(*stack) - 1
		var top_element T = (*stack)[index]
		return errors.New("NullPointerException"), top_element
	} else {
		var index int = len(*stack) - 1
		var top_element T = (*stack)[index]
		return nil, top_element
	}
}

func CalculatorHelper(operators Stack_Implementation[rune], operands Stack_Implementation[int]) {
	if operators.IsEmpty() {
		return
	}
	_, right_operand := operands.Top()
	operands.Pop()
	_, left_operand := operands.Top()
	operands.Pop()
	_, operator := operators.Top()
	operators.Pop()

	// UTF-8 operators
	// '*' = 42. hex = 0x2a
	// '+' = 43. hex = 0x2b
	// '-' = 45. hex = 0x2d
	// '/' = 47. hex = 0x2f

	if operator == 42 {
		operands.InsertIn(left_operand * right_operand)
	} else if operator == 47 {
		operands.InsertIn(left_operand / right_operand)
	} else if operator == 43 {
		operands.InsertIn(left_operand + right_operand)
	} else if operator == 45 {
		operands.InsertIn(left_operand - right_operand)
	}
}

func IsDigit(unicode byte) bool {
	var byte_decimal []byte = []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	for i := 0; i < len(byte_decimal); i++ {
		if byte_decimal[i] == unicode {
			return true
		}
	}
	return false
}

package algorithms_and_data_structures

import "testing"

func TestReverseString(t *testing.T) {
	result := "hello"
	test := ReverseString("hello")
	if result != test {
		t.Errorf("%t", false)
	}
}

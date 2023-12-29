package algorithms_and_data_structures

import "testing"

func TestFindNumber(t *testing.T) {
	var array_var = [10]int{1, 2, 3, 4}
	vbool := FindNumber(array_var, 5)
	if vbool != true {
		t.Errorf("FindNumber(array_var, target) = %t; correct = true", vbool)
	}
}

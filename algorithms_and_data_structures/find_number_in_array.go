package algorithms_and_data_structures

func FindNumber(numbers [10]int, target int) bool {
	// var index int
	// var val int
	var isTrue bool = true
	if target == 0 {
		isTrue = false
	}
	for _, k := range numbers {
		if numbers[k] != target {
			// index = i
			// val = k
			isTrue = false
		} else {
			isTrue = true
		}
	}
	return isTrue
}

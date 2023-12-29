package algorithms_and_data_structures

import "fmt"

func MaxSubArray(array []int, n int) int {
	if n > len(array) {
		return 0
	}

	maxSum := 0
	tempSum := 0
	for i := 0; i < n; i++ {
		maxSum += array[i]
	}

	tempSum = maxSum

	for i := n; i < len(array); i++ {
		fmt.Println(tempSum, array[i-n], array[i])
		tempSum = tempSum - array[i-n] + array[i]
		// fmt.Println(tempSum)
		maxSum = max(maxSum, tempSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

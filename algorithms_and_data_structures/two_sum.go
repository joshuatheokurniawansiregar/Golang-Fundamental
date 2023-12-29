package algorithms_and_data_structures

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
// You can return the answer in any order.

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

//Using HashMap
func TwoSum(nums []int, target int) []int {
	var maps = make(map[int]int)
	var array_temp = []int{}
	for i, each := range nums {
		if pos, ok := maps[target-each]; ok {
			array_temp = []int{pos, i}
			return array_temp
		}
		maps[each] = i
	}
	return []int{}
}

func TwoSumWithTwoPointer(array []int, target int) []int {
	var res []int = []int{}
	var ptr1, ptr2 int = 0, len(array) - 1
	for ptr1 < ptr2 {
		var curr int = array[ptr2] + array[ptr1]
		if curr == target {
			res = []int{ptr1, ptr2}
			return res
		} else if curr > target {
			ptr2--
		} else {
			ptr1++
		}
	}
	return []int{}
}

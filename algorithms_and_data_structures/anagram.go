package algorithms_and_data_structures

import (
	"sort"
)

func Anagram(s, t string) bool {
	var isAnagram bool = true
	var chars_s []byte = []byte(s)
	var chars_t []byte = []byte(s)
	sort.Slice(chars_s, func(i, j int) bool {
		return chars_s[i] < chars_t[j]
	})
	sort.Slice(chars_t, func(i, j int) bool {
		return chars_t[i] < chars_t[j]
	})
	var string_s = string(chars_s)
	var string_t = string(chars_t)
	var length_s = len(string_s)
	var length_t = len(string_t)

	if length_s != length_t {
		isAnagram = false
	}
	for i := 0; i < length_s; i++ {
		if chars_s[i] != chars_t[i] {
			isAnagram = false
		}
	}
	return isAnagram
}

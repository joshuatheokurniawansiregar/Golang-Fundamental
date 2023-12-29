package algorithms_and_data_structures

func ReverseString(word string) string {
	var j int = 0
	var reversed_word = make([]byte, len(word))
	for i := len(word) - 1; i >= 0; i-- {
		var one_char byte = word[i]
		reversed_word[j] = one_char
		j++
	}
	word = string(reversed_word)
	return word
}

package algorithms_and_data_structures

var words string = "hello world hello world. Hello world. Nama saya Joshua Theo Kurniawan. Saya bukan orang ganteng."

func ParsingStringPattern(text string, pattern string) (int, int) {
	x, y := -1, -1
	var text_length uint64 = uint64(len(text))
	var pattern_length uint64 = uint64(len(pattern))

	for i := 0; i <= int(text_length)-int(pattern_length); i++ {
		var j int
		for j = 0; j < int(pattern_length); j++ {
			if text[i+j] != pattern[j] {
				break
			}
		}
		if j == int(pattern_length) {
			x, y = i, i+1
		}

	}
	return x, y
}

func SplitSubString(text string, pattern_punc string) string {
	// var res_text []string = make([]string, len(text))
	var text_length uint64 = uint64(len(text))
	var pattern_length uint64 = uint64(len(pattern_punc))
	var word string

	for i := 0; i <= int(text_length)-int(pattern_length); i++ {
		var j int
		for j = 0; j < int(pattern_length); j++ {
			if text[i+j] != pattern_punc[j] {
				break
			}
		}
		if j != int(pattern_length) {
			word = word + string(text[i])
		}
	}
	return word
}

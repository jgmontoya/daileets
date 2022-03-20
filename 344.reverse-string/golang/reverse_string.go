package reverse_string

func ReverseString(s []byte) {
	left_index := 0
	right_index := len(s) - 1

	for left_index < right_index {
		swap(s, left_index, right_index)
		left_index++
		right_index--
	}
}

func swap(s []byte, first int, second int) {
	mem := s[first]
	s[first] = s[second]
	s[second] = mem
}

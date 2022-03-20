package reverse_words

func ReverseWords(s string) string {
	reversed_words := make([]byte, len(s))
	index := 0
	last_index_of_word := last_char_of_word_index(s, index)

	for last_index_of_word < len(s) {
		insert_word(s, index, last_index_of_word, reversed_words)
		index = last_index_of_word + 1
		if index < len(s) {
			reversed_words[index] = ' '
		}
		index++
		last_index_of_word = last_char_of_word_index(s, index)
	}
	return string(reversed_words)
}

func swap_insert(s string, first int, second int, reversed_words []byte) {
	reversed_words[first] = s[second]
	reversed_words[second] = s[first]
}

func last_char_of_word_index(s string, index int) int {
	for index < len(s) && s[index] != ' ' {
		index++
	}
	return index - 1
}

func insert_word(s string, index int, last_index int, reverse_words []byte) {
	for index <= last_index {
		swap_insert(s, index, last_index, reverse_words)
		index++
		last_index--
	}
}

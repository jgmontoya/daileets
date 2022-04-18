package first_uniq_char

func firstUniqChar(s string) int {
	char_repeat := map[rune]int{}

	for _, char := range s {
		char_repeat[char]++
	}

	for index, char := range s {
		if char_repeat[char] == 1 {
			return index
		}
	}
	return -1
}

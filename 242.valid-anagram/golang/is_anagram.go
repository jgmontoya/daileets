package is_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_letters := map[rune]int{}
	for _, char := range s {
		s_letters[char]++
	}

	for _, char := range t {
		s_letters[char]--
		if s_letters[char] < 0 {
			return false
		}
	}
	return true
}

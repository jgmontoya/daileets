package can_construct

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	ransom_letters := map[rune]int{}
	for _, char := range ransomNote {
		ransom_letters[char]++
	}

	used_letters := 0
	for _, char := range magazine {
		if ransom_letters[char] > 0 {
			used_letters++
			ransom_letters[char]--
		}
		if used_letters == len(ransomNote) {
			return true
		}
	}
	return false
}

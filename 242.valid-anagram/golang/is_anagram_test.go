package is_anagram

import "testing"

func CheckIsAnagram(t *testing.T, word1 string, word2 string, solution bool) {
	is_anagram := isAnagram(word1, word2)
	if is_anagram == solution {
		t.Log("OK")
	} else {
		t.Errorf("isAnagram(%s, %s): Got %t; expected %t", word1, word2, is_anagram, solution)
		t.Fail()
	}
}

func TestIsAnagram(t *testing.T) {
	word1, word2 := "anagram", "nagaram"
	solution := true
	CheckIsAnagram(t, word1, word2, solution)

	word1, word2 = "rat", "car"
	solution = false
	CheckIsAnagram(t, word1, word2, solution)
}

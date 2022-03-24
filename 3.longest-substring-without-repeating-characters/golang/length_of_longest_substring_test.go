package length_of_longest_substring

import "testing"

func CheckLengthOfLongestSubstring(t *testing.T, input string, solution int) {
	length_of_longest_substring := LengthOfLongestSubstring(input)
	if length_of_longest_substring == solution {
		t.Logf("LengthOfLongestSubstring(%s) = %d: OK", input, solution)
	} else {
		t.Errorf("LengthOfLongestSubstring(%s): Got %d; expected %d", input, length_of_longest_substring, solution)
		t.Fail()
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	input := "abcabcbb"
	solution := 3
	CheckLengthOfLongestSubstring(t, input, solution)

	input = "bbbbb"
	solution = 1
	CheckLengthOfLongestSubstring(t, input, solution)

	input = "pwwkew"
	solution = 3
	CheckLengthOfLongestSubstring(t, input, solution)
}

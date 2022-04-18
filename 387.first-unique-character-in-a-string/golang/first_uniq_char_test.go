package first_uniq_char

import "testing"

func CheckFirstUniqChar(t *testing.T, input string, solution int) {
	first_uniq_char := firstUniqChar(input)
	if first_uniq_char == solution {
		t.Log("OK")
	} else {
		t.Errorf("firstUniqChar(%s): Got %d; expected %d", input, first_uniq_char, solution)
		t.Fail()
	}
}

func TestFirstUniqChar(t *testing.T) {
	input := "leetcode"
	solution := 0
	CheckFirstUniqChar(t, input, solution)

	input = "loveleetcode"
	solution = 2
	CheckFirstUniqChar(t, input, solution)

	input = "aabb"
	solution = -1
	CheckFirstUniqChar(t, input, solution)
}

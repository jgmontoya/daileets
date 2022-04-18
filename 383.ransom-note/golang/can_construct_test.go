package can_construct

import "testing"

func CheckCanConstruct(t *testing.T, ransomNote string, magazine string, solution bool) {
	can_construct := canConstruct(ransomNote, magazine)
	if can_construct == solution {
		t.Log("OK")
	} else {
		t.Errorf("canConstruct(%s, %s): Got %t; expected %t", ransomNote, magazine, can_construct, solution)
		t.Fail()
	}
}
func TestCanConstruct(t *testing.T) {
	ransomNote := "a"
	magazine := "b"
	solution := false
	CheckCanConstruct(t, ransomNote, magazine, solution)

	ransomNote = "aa"
	magazine = "ab"
	solution = false
	CheckCanConstruct(t, ransomNote, magazine, solution)

	ransomNote = "aa"
	magazine = "aab"
	solution = true
	CheckCanConstruct(t, ransomNote, magazine, solution)
}

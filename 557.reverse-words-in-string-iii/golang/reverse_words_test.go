package reverse_words

import (
	"testing"
)

func CheckReverseWords(t *testing.T, input string, solution string) {
	reversed_words := ReverseWords(input)

	if reversed_words == solution {
		t.Logf("OK")
	} else {
		t.Errorf("Got %s; expected %s", reversed_words, solution)
		t.Fail()
	}

}

func TestReverseWords(t *testing.T) {
	input := "Let's take LeetCode contest"
	solution := "s'teL ekat edoCteeL tsetnoc"
	CheckReverseWords(t, input, solution)
	input = "God Ding"
	solution = "doG gniD"
	CheckReverseWords(t, input, solution)
}

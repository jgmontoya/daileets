package reverse_string

import (
	"bytes"
	"testing"
)

func CheckReverseString(t *testing.T, input []byte, solution []byte) {
	ReverseString(input)

	if bytes.Equal(input, solution) {
		t.Logf("OK")
	} else {
		t.Errorf("Got %s; expected %s", input, solution)
		t.Fail()
	}
}

func TestReverseString(t *testing.T) {
	input := []byte{'h', 'e', 'l', 'l', 'o'}
	solution := []byte{'o', 'l', 'l', 'e', 'h'}
	CheckReverseString(t, input, solution)

	input = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	solution = []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	CheckReverseString(t, input, solution)
}

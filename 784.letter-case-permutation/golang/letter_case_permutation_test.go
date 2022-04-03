package letter_case_permutation

import (
	"testing"
)

func SliceEqual(first_slice []string, second_slice []string) bool {
	first_length := len(first_slice)
	second_length := len(second_slice)

	if first_length != second_length {
		return false
	}

	for i := 0; i < first_length; i++ {
		if first_slice[i] != second_slice[i] {
			return false
		}
	}
	return true
}

func CheckletterCasePermutation(t *testing.T, input string, solution []string) {
	// This check should be done using sets as order does not matter.
	// This only works if one has a known order for the output
	combinations := letterCasePermutation(input)
	if SliceEqual(combinations, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %s; expected %s", combinations, solution)
		t.Fail()
	}
}

func TestLetterCasePermutation(t *testing.T) {
	input := "a1b2"
	solution := []string{"a1b2", "A1b2", "a1B2", "A1B2"}
	CheckletterCasePermutation(t, input, solution)

	input = "3z4"
	solution = []string{"3z4", "3Z4"}
	CheckletterCasePermutation(t, input, solution)

	input = "3"
	solution = []string{"3"}
	CheckletterCasePermutation(t, input, solution)

	input = "z"
	solution = []string{"z", "Z"}
	CheckletterCasePermutation(t, input, solution)
}

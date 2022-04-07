package contains_duplicate

import "testing"

func CheckContainsDuplicate(t *testing.T, input []int, solution bool) {
	contains_duplicate := containsDuplicate(input)
	if contains_duplicate == solution {
		t.Log("OK")
	} else {
		t.Errorf("contains_duplicate(%d): Got %t; expected %t", input, contains_duplicate, solution)
		t.Fail()
	}
}

func TestContainsDuplicate(t *testing.T) {
	input := []int{1, 2, 3, 1}
	solution := true
	CheckContainsDuplicate(t, input, solution)

	input = []int{1, 2, 3, 4}
	solution = false
	CheckContainsDuplicate(t, input, solution)

	input = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	solution = true
	CheckContainsDuplicate(t, input, solution)
}

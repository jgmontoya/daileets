package permute

import "testing"

func SliceEqual(first_slice []int, second_slice []int) bool {
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

func SliceEqual2D(first_slice [][]int, second_slice [][]int) bool {
	first_length := len(first_slice)
	second_length := len(second_slice)

	if first_length != second_length {
		return false
	}

	for i := 0; i < first_length; i++ {
		if !SliceEqual(first_slice[i], second_slice[i]) {
			return false
		}
	}
	return true
}

func Checkpermute(t *testing.T, input []int, solution [][]int) {
	// This check should be done using sets as order does not matter.
	// This only works if one has a known order for the output
	combinations := permute(input)
	if SliceEqual2D(combinations, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", combinations, solution)
		t.Fail()
	}
}
func TestCombine(t *testing.T) {
	input := []int{1, 2, 3}
	solution := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2},
		{3, 2, 1}}
	Checkpermute(t, input, solution)

	input = []int{0, 1}
	solution = [][]int{{0, 1}, {1, 0}}
	Checkpermute(t, input, solution)

	input = []int{1}
	solution = [][]int{{1}}
	Checkpermute(t, input, solution)
}

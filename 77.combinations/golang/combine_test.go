package combine

import (
	"testing"
)

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

func Checkcombine(t *testing.T, n int, k int, solution [][]int) {
	// This check should be done using sets as order does not matter.
	// This only works if one has a known order for the output
	combinations := combine(n, k)
	if SliceEqual2D(combinations, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", combinations, solution)
		t.Fail()
	}
}
func TestCombine(t *testing.T) {
	n := 4
	k := 2
	solution := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	Checkcombine(t, n, k, solution)

	n = 2
	k = 1
	solution = [][]int{{1}, {2}}
	Checkcombine(t, n, k, solution)

	n = 4
	k = 3
	solution = [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}
	Checkcombine(t, n, k, solution)

}

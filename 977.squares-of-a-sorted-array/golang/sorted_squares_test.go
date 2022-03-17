package sorted_squares

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

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	solution := []int{0, 1, 9, 16, 100}

	sorted_squares := SortedSquares(nums)
	if SliceEqual(sorted_squares, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", sorted_squares, solution)
		t.Fail()
	}

	nums = []int{-7, -3, 2, 3, 11}
	solution = []int{4, 9, 9, 49, 121}

	sorted_squares = SortedSquares(nums)
	if SliceEqual(sorted_squares, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", sorted_squares, solution)
		t.Fail()
	}
}

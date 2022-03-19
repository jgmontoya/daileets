package move_zeroes

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

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	solution := []int{1, 3, 12, 0, 0}

	MoveZeroes(nums)
	if SliceEqual(nums, solution) {
		t.Logf("OK")
	} else {
		t.Errorf("Got %d; expected %d", nums, solution)
		t.Fail()
	}

	nums = []int{0}
	solution = []int{0}
	MoveZeroes(nums)
	if SliceEqual(nums, solution) {
		t.Logf("OK")
	} else {
		t.Errorf("Got %d; expected %d", nums, solution)
		t.Fail()
	}
}

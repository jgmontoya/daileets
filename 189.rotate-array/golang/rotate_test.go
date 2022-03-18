package rotate

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

func CheckRotation(t *testing.T, nums []int, k int, solution []int) {
	Rotate(nums, k)
	if SliceEqual(nums, solution) {
		t.Logf("Rotate(%d, %d): OK", nums, k)
	} else {
		t.Errorf("Rotate(%d, %d): Got %d; expected %d", nums, k, nums, solution)
		t.Fail()
	}
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	CheckRotation(t, nums, 0, []int{1, 2, 3, 4, 5, 6, 7})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 1, []int{7, 1, 2, 3, 4, 5, 6})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 2, []int{6, 7, 1, 2, 3, 4, 5})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 3, []int{5, 6, 7, 1, 2, 3, 4})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 4, []int{4, 5, 6, 7, 1, 2, 3})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 5, []int{3, 4, 5, 6, 7, 1, 2})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 6, []int{2, 3, 4, 5, 6, 7, 1})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 7, []int{1, 2, 3, 4, 5, 6, 7})
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	CheckRotation(t, nums, 8, []int{7, 1, 2, 3, 4, 5, 6})
}

package two_sum

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

func CheckTwoSum(t *testing.T, nums []int, target int, solution []int) {
	two_sum := twoSum(nums, target)
	if SliceEqual(two_sum, solution) {
		t.Log("OK")
	} else {
		t.Errorf("twoSum(%d, %d): Got %d; expected %d", nums, target, two_sum, solution)
		t.Fail()
	}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	solution := []int{0, 1}
	CheckTwoSum(t, nums, target, solution)

	nums = []int{3, 2, 4}
	target = 6
	solution = []int{1, 2}
	CheckTwoSum(t, nums, target, solution)

	nums = []int{3, 3}
	target = 6
	solution = []int{0, 1}
	CheckTwoSum(t, nums, target, solution)
}

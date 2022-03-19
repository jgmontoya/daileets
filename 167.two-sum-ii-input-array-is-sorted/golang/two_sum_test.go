package two_sum

import "testing"

func CheckTwoSum(t *testing.T, numbers []int, target int, solution []int) {
	two_sum := TwoSum(numbers, target)
	if two_sum[0] == solution[0] && two_sum[1] == solution[1] {
		t.Logf("TwoSum(%d, %d): OK", numbers, target)
	} else {
		t.Errorf("TwoSum(%d, %d): Got %d; expected %d", numbers, target, two_sum, solution)
		t.Fail()
	}
}

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	solution := []int{1, 2}
	CheckTwoSum(t, numbers, target, solution)

	numbers = []int{2, 3, 4}
	target = 6
	solution = []int{1, 3}
	CheckTwoSum(t, numbers, target, solution)

	numbers = []int{-1, 0}
	target = -1
	solution = []int{1, 2}
	CheckTwoSum(t, numbers, target, solution)

}

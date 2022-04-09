package max_sub_array

import "testing"

func CheckMaxSubArray(t *testing.T, input []int, solution int) {
	max_sub_array := maxSubArray(input)
	if max_sub_array == solution {
		t.Log("OK")
	} else {
		t.Errorf("maxSubArray(%d); Got %d, expected %d", input, max_sub_array, solution)
	}
}

func TestMaxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	solution := 6
	CheckMaxSubArray(t, input, solution)

	input = []int{1}
	solution = 1
	CheckMaxSubArray(t, input, solution)

	input = []int{5, 4, -1, 7, 8}
	solution = 23
	CheckMaxSubArray(t, input, solution)

	input = []int{-2, 1}
	solution = 1
	CheckMaxSubArray(t, input, solution)
}

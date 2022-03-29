package update_matrix

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

func CheckUpdateMatrix(t *testing.T, input [][]int, solution [][]int) {
	updated_matrix := UpdateMatrix(input)
	if SliceEqual2D(updated_matrix, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", updated_matrix, solution)
		t.Fail()
	}
}

func TestUpdateMatrix(t *testing.T) {
	input := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	solution := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	CheckUpdateMatrix(t, input, solution)

	input = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	solution = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}
	CheckUpdateMatrix(t, input, solution)

	input = [][]int{
		{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
		{1, 1, 1, 1, 0, 1, 0, 0, 1, 1},
	}
	solution = [][]int{
		{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 2, 1, 1, 0, 1},
		{2, 1, 1, 1, 1, 2, 1, 0, 1, 0},
		{3, 2, 2, 1, 0, 1, 0, 0, 1, 1},
	}
	CheckUpdateMatrix(t, input, solution)

	input = [][]int{{1}, {1}, {1}, {0}, {1}, {1}, {1}, {0}}
	solution = [][]int{{3}, {2}, {1}, {0}, {1}, {2}, {1}, {0}}
	CheckUpdateMatrix(t, input, solution)
}

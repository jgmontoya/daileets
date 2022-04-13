package generate

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

func CheckGenerate(t *testing.T, input int, solution [][]int) {
	triangle := generate(input)
	if SliceEqual2D(triangle, solution) {
		t.Log("OK")
	} else {
		t.Errorf("generate(%d): Got %d; expected %d", input, triangle, solution)
	}
}

func TestGenerate(t *testing.T) {
	input := 5
	solution := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	CheckGenerate(t, input, solution)

	input = 1
	solution = [][]int{{1}}
	CheckGenerate(t, input, solution)
}

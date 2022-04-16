package search_matrix

import "testing"

func CheckSearchMatrix(t *testing.T, matrix [][]int, target int, solution bool) {
	search_matrix := searchMatrix(matrix, target)
	if search_matrix == solution {
		t.Log("OK")
	} else {
		t.Errorf("searchMatrix(%d, %d): Got %t; expected %t", matrix, target, search_matrix, solution)
		t.Fail()
	}
}
func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	solution := true
	// CheckSearchMatrix(t, matrix, target, solution)

	target = 13
	solution = false
	// CheckSearchMatrix(t, matrix, target, solution)

	matrix = [][]int{{1}}
	target = 2
	solution = false
	CheckSearchMatrix(t, matrix, target, solution)
}

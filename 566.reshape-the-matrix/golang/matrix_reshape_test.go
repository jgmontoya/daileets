package matrix_reshape

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

func CheckMatrixReshape(t *testing.T, mat [][]int, r int, c int, solution [][]int) {
	matrix_reshape := matrixReshape(mat, r, c)
	if SliceEqual2D(matrix_reshape, solution) {
		t.Log("OK")
	} else {
		t.Errorf("matrixReshape(%d, %d, %d): Got %d; expected %d", mat, r, c, matrix_reshape, solution)
		t.Fail()
	}
}

func TestMatrixReshape(t *testing.T) {
	mat := [][]int{{1, 2}, {3, 4}}
	r, c := 1, 4
	solution := [][]int{{1, 2, 3, 4}}
	CheckMatrixReshape(t, mat, r, c, solution)

	mat = [][]int{{1, 2}, {3, 4}}
	r, c = 2, 4
	solution = [][]int{{1, 2}, {3, 4}}
	CheckMatrixReshape(t, mat, r, c, solution)
}

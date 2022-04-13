package matrix_reshape

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	result := make([][]int, r)
	for index := 0; index < r; index++ {
		result[index] = make([]int, c)
	}
	element_index := 0
	for _, row := range mat {
		for _, num := range row {
			new_row := element_index / c
			new_col := element_index % c
			result[new_row][new_col] = num
			element_index++
		}
	}
	return result
}

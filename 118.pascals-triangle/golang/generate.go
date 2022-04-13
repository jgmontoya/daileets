package generate

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}

	for row := 1; row < numRows; row++ {
		new_row := make([]int, row+1)
		new_row[0] = 1
		new_row[row] = 1
		for col_index := 1; col_index < row; col_index++ {
			new_row[col_index] = result[row-1][col_index-1] + result[row-1][col_index]
		}
		result[row] = new_row
	}

	return result
}

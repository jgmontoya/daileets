package update_matrix

func UpdateMatrix(mat [][]int) [][]int {
	for row_index, row := range mat {
		for col_index, element := range row {
			if element != 0 {
				dist := len(mat) + len((mat)[0])
				if row_index > 0 {
					dist = Min(dist, mat[row_index-1][col_index]+1)
				}
				if col_index > 0 {
					dist = Min(dist, mat[row_index][col_index-1]+1)
				}
				mat[row_index][col_index] = dist
			}
		}
	}

	for row_index := len(mat) - 1; row_index >= 0; row_index-- {
		for col_index := len(mat[0]) - 1; col_index >= 0; col_index-- {
			dist := mat[row_index][col_index]
			if row_index < len(mat)-1 {
				dist = Min(dist, mat[row_index+1][col_index]+1)
			}
			if col_index < len(mat[0])-1 {
				dist = Min(dist, mat[row_index][col_index+1]+1)
			}
			mat[row_index][col_index] = dist
		}
	}
	return mat
}

func Min(num1 int, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}

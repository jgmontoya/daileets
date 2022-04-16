package search_matrix

func searchMatrix(matrix [][]int, target int) bool {
	num_rows := len(matrix)
	num_cols := len(matrix[0])
	return recursive_search_matrix(matrix, target, 0, num_rows*num_cols)
}

func recursive_search_matrix(matrix [][]int, target int, from_index int, to_index int) bool {
	if to_index < from_index {
		return false
	}
	mid_index := from_index + (to_index-from_index)/2
	if mid_index > len(matrix)*len(matrix[0])-1 {
		return false
	}
	mid_element := linear_to_matrix_indexer(matrix, mid_index)
	if target == mid_element {
		return true
	} else if target > mid_element {
		return recursive_search_matrix(matrix, target, mid_index+1, to_index)
	} else {
		return recursive_search_matrix(matrix, target, from_index, mid_index-1)
	}
}

func linear_to_matrix_indexer(matrix [][]int, index int) int {
	num_cols := len(matrix[0])
	row := index / num_cols
	col := index % num_cols
	return matrix[row][col]
}

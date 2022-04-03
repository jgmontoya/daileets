package minimum_total

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	smaller_triangle := reduceLastRow(triangle)
	for len(smaller_triangle) > 1 {
		smaller_triangle = reduceLastRow(smaller_triangle)
	}
	return smaller_triangle[0][0]
}

func reduceLastRow(triangle [][]int) [][]int {
	last_index := len(triangle) - 1
	second_to_last_index := last_index - 1

	for col_index, _ := range triangle[second_to_last_index] {
		left_son := triangle[last_index][col_index]
		right_son := triangle[last_index][col_index+1]
		triangle[second_to_last_index][col_index] += min(left_son, right_son)
	}
	return triangle[:last_index]
}

func min(i int, j int) int {
	if i <= j {
		return i
	}
	return j
}

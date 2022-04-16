package is_valid_sudoku

func isValidSudoku(board [][]byte) bool {
	for index := 0; index < 9; index++ {
		if faulty_row(board, index) || faulty_col(board, index) || faulty_square(board, index) {
			return false
		}
	}
	return true
}

func faulty_row(board [][]byte, index int) bool {
	row := board[index]
	nums := map[byte]bool{}
	for _, num := range row {
		if num == '.' {
			continue
		}
		if nums[num] {
			return true
		}
		nums[num] = true
	}
	return false
}

func faulty_col(board [][]byte, index int) bool {
	nums := map[byte]bool{}
	for row_index := 0; row_index < 9; row_index++ {
		num := board[row_index][index]
		if num == '.' {
			continue
		}
		if nums[num] {
			return true
		}
		nums[num] = true
	}
	return false
}

func faulty_square(board [][]byte, index int) bool {
	nums := map[byte]bool{}
	for row_index := (index % 3) * 3; row_index < (index%3)*3+3; row_index++ {
		for col_index := (index / 3) * 3; col_index < (index/3)*3+3; col_index++ {
			num := board[row_index][col_index]
			if num == '.' {
				continue
			}
			if nums[num] {
				return true
			}
			nums[num] = true
		}
	}
	return false
}

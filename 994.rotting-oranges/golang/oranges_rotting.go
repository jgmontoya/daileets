package oranges_rotting

func OrangesRotting(grid [][]int) int {
	minutes_elapsed := 0
	rotten_to_propagate := OrangeQueue{}
	fresh_to_rott := 0

	for row_index, row := range grid {
		for col_index, orange_state := range row {
			if orange_state == 2 {
				rotten_to_propagate.Push(Orange{row: row_index, col: col_index,
					state: orange_state, minute_rotten: 0,
				})
			} else if orange_state == 1 {
				fresh_to_rott++
			}
		}
	}

	if rotten_to_propagate.Length() == 0 && fresh_to_rott != 0 {
		return -1
	}

	if fresh_to_rott == 0 {
		return 0
	}

	for rotten_to_propagate.Length() != 0 {
		rotten_orange := rotten_to_propagate.Pop()
		if rotten_orange.minute_rotten > minutes_elapsed {
			minutes_elapsed = rotten_orange.minute_rotten
		}
		for _, neighbor := range Neighbors(&grid, rotten_orange) {
			if neighbor.state == 1 {
				neighbor.state = 2
				grid[neighbor.row][neighbor.col] = 2
				neighbor.minute_rotten = rotten_orange.minute_rotten + 1
				rotten_to_propagate.Push(neighbor)
				fresh_to_rott--
			}
		}
	}

	if fresh_to_rott == 0 {
		return minutes_elapsed
	}

	return -1
}

func Neighbors(grid *[][]int, orange Orange) []Orange {
	row_max := len((*grid))
	col_max := len((*grid)[0])
	neighbors := []Orange{}
	for _, shift := range []int{-1, 1} {
		west_east_neighbor_col := orange.col + shift
		north_south_neighbor_row := orange.row + shift
		if west_east_neighbor_col >= 0 && west_east_neighbor_col < col_max {
			neighbor_orange_row := orange.row
			neighbor_orange := Orange{row: orange.row, col: west_east_neighbor_col,
				state: (*grid)[neighbor_orange_row][west_east_neighbor_col],
			}
			neighbors = append(neighbors, neighbor_orange)
		}

		if north_south_neighbor_row >= 0 && north_south_neighbor_row < row_max {
			neighbor_orange_col := orange.col
			neighbor_orange := Orange{row: north_south_neighbor_row, col: orange.col,
				state: (*grid)[north_south_neighbor_row][neighbor_orange_col]}
			neighbors = append(neighbors, neighbor_orange)
		}
	}
	return neighbors
}

type Orange struct {
	row           int
	col           int
	state         int
	minute_rotten int
}

type OrangeQueue struct {
	container []Orange
}

func (oq *OrangeQueue) Push(orange Orange) {
	oq.container = append(oq.container, orange)
}

func (oq *OrangeQueue) Pop() Orange {
	orange := oq.container[0]
	oq.container = oq.container[1:]
	return orange
}

func (oq *OrangeQueue) Length() int {
	return len(oq.container)
}

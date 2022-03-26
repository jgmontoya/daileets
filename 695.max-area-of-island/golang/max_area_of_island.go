package max_area_of_island

func MaxAreaOfIsland(grid [][]int) int {
	max_area_of_island := 0

	for row_index, row := range grid {
		for col_index, spot := range row {
			if spot != 1 {
				continue
			}

			area_of_island := AreaOfIsland(&grid, row_index, col_index)
			if area_of_island > max_area_of_island {
				max_area_of_island = area_of_island
			}
		}
	}

	return max_area_of_island
}

func AreaOfIsland(grid *[][]int, row_index int, col_index int) int {
	stack := PointStack{}
	stack.Push(Point{row: row_index, col: col_index})
	area := 0

	for stack.Length() > 0 {
		point_to_paint := stack.Pop()
		if (*grid)[point_to_paint.row][point_to_paint.col] == 1 {
			area++
			(*grid)[point_to_paint.row][point_to_paint.col] = 2
		}

		for _, neighbor := range Neighbors(grid, point_to_paint) {
			neighbor_is_island := (*grid)[neighbor.row][neighbor.col]
			if neighbor_is_island == 1 {
				stack.Push(neighbor)
			}
		}
	}

	return area
}

func Neighbors(grid *[][]int, point Point) []Point {
	row_max := len((*grid))
	col_max := len((*grid)[0])
	neighbors := []Point{}
	for _, shift := range []int{-1, 1} {
		west_east_neighbor_col := point.col + shift
		north_south_neighbor_row := point.row + shift
		if west_east_neighbor_col >= 0 && west_east_neighbor_col < col_max {
			neighbors = append(neighbors, Point{row: point.row, col: west_east_neighbor_col})
		}

		if north_south_neighbor_row >= 0 && north_south_neighbor_row < row_max {
			neighbors = append(neighbors, Point{row: north_south_neighbor_row, col: point.col})
		}
	}
	return neighbors
}

type Point struct {
	row int
	col int
}

type PointStack struct {
	container []Point
}

func (pq *PointStack) Push(point Point) {
	pq.container = append(pq.container, point)
}

func (pq *PointStack) Pop() Point {
	last_index := len(pq.container) - 1
	point := pq.container[last_index]
	pq.container = pq.container[:last_index]
	return point
}

func (pq *PointStack) Length() int {
	return len(pq.container)
}

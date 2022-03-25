package flood_fill

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	paint_queue := PointQueue{}
	paint_queue.Push(Point{row: sr, col: sc})
	original_color := image[sr][sc]
	if newColor == original_color {
		return image
	}
	for paint_queue.Length() > 0 {
		PopAndPaint(&paint_queue, &image, original_color, newColor)
	}
	return image
}

func PopAndPaint(paint_queue *PointQueue, image *[][]int, original_color int, color int) {
	point_to_paint := (*paint_queue).Pop()
	(*image)[point_to_paint.row][point_to_paint.col] = color

	for _, neighbor := range Neighbors(image, point_to_paint) {
		neighbor_color := (*image)[neighbor.row][neighbor.col]
		if neighbor_color == original_color {
			(*paint_queue).Push(neighbor)
		}
	}
}

func Neighbors(image *[][]int, point Point) []Point {
	row_max := len((*image))
	col_max := len((*image)[0])
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

type PointQueue struct {
	container []Point
}

func (pq *PointQueue) Pop() Point {
	point := pq.container[0]
	pq.container = pq.container[1:]
	return point
}

func (pq *PointQueue) Push(point Point) {
	pq.container = append(pq.container, point)
}

func (pq *PointQueue) Length() int {
	return len(pq.container)
}

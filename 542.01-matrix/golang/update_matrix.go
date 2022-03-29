package update_matrix

func UpdateMatrix(mat [][]int) [][]int {
	for row_index, row := range mat {
		for col_index, element := range row {
			if element != 0 {
				point := Point{row: row_index, col: col_index}
				mat[point.row][point.col] = Dist_to_0(&mat, point)
			}
		}
	}
	return mat
}

func Dist_to_0(mat *[][]int, point Point) int {
	dist := len(*mat) + len((*mat)[0])
	if point.row > 0 {
		dist = Min(dist, (*mat)[point.row-1][point.col]+1)
	}
	if point.col > 0 {
		dist = Min(dist, (*mat)[point.row][point.col-1]+1)
	}
	found_0 := false
	for row_index := point.row; row_index < len(*mat); row_index++ {
		for col_index := Max(0, point.col-dist+1); col_index < len((*mat)[0]); col_index++ {
			if row_index > point.row+dist {
				found_0 = true
				break
			}
			if col_index > point.col+dist {
				break
			}
			checking := Point{row: row_index, col: col_index}
			if (*mat)[row_index][col_index] == 0 {
				dist_to_checking := ManhattanDist(point, checking)
				dist = Min(dist, dist_to_checking)
			}
		}
		if found_0 {
			break
		}
	}
	return dist
}

func Min(num1 int, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}

func Max(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}

func Abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func ManhattanDist(point1 Point, point2 Point) int {
	return Abs(point1.row-point2.row) + Abs(point1.col-point2.col)
}

type Point struct {
	row int
	col int
}

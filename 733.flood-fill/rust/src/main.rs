use std::collections::VecDeque;

fn main() {}

struct Solution {}

struct Point {
    row: usize,
    col: usize,
}

impl Solution {
    pub fn flood_fill(image: Vec<Vec<i32>>, sr: i32, sc: i32, color: i32) -> Vec<Vec<i32>> {
        let original_color = image[sr as usize][sc as usize];
        if color == original_color {
            return image;
        }
        let mut image_clone = image.clone();
        let mut paint_queue: VecDeque<Point> = VecDeque::new();
        paint_queue.push_back(Point{row: sr as usize, col: sc as usize});
        while paint_queue.len() > 0 {
            Solution::pop_and_paint(&mut paint_queue, &mut image_clone, original_color, color);
        }
        return image_clone
    }

    fn pop_and_paint(paint_queue: &mut VecDeque<Point>, image: &mut Vec<Vec<i32>>, original_color: i32, new_color: i32) {

        let mut point_to_paint = paint_queue.pop_front().unwrap();
        image[point_to_paint.row as usize][point_to_paint.col as usize] = new_color;

        for neighbor in Solution::neighbors(image, point_to_paint) {
            let neighbor_color = image[neighbor.row][neighbor.col];
            if neighbor_color == original_color {
                paint_queue.push_back(neighbor)
            }
        }
    }

    fn neighbors(image: &mut Vec<Vec<i32>>, point: Point) -> Vec<Point> {
        let row_max = image.len();
        let col_max = image[0].len();
        let mut neighbors: Vec<Point> = Vec::new();

        for shift in vec![-1, 1] {
            let west_east_neighbor_col =( point.col as i32 + shift) as usize;
            let north_south_neighbor_row = (point.row as i32 + shift) as usize;

            if west_east_neighbor_col >= 0 && west_east_neighbor_col < col_max {
                neighbors.push(Point{row: point.row, col: west_east_neighbor_col});
            }

            if north_south_neighbor_row >= 0 && north_south_neighbor_row < row_max {
                neighbors.push(Point{row: north_south_neighbor_row, col: point.col});
            }
        }

        return neighbors;
    }
}

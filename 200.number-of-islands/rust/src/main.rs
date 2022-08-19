struct Solution {}

fn main() {}

struct Point {
    row: i32,
    col: i32,
}

impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let mut num_islands = 0;

        let num_rows = grid.len();
        let num_cols = grid[0].len();

        let mut map = grid.clone();

        for row in 0..num_rows as i32 {
            for col in 0..num_cols as i32 {
                if map[row as usize][col as usize] == '1' {
                    num_islands += 1;
                    Solution::clear_island(&mut map, Point{row, col});
                }
            }
        }
        return num_islands
    }

    fn clear_island(map: &mut Vec<Vec<char>>, point: Point) {
        if map[point.row as usize][point.col as usize] == '0' { return }

        map[point.row as usize][point.col as usize] = '0';
        for neighbor in Solution::neighbors(map, point) {
            Solution::clear_island(map, neighbor);
        }
    }

    fn neighbors(grid: &Vec<Vec<char>>, point: Point) -> Vec<Point> {
        let row_max = grid.len();
        let col_max = grid[0].len();
        let mut neighbors: Vec<Point> = Vec::with_capacity(4);

        for shift in vec![-1, 1] {
            let west_east_neighbor_col = point.col as i32 + shift;
            let north_south_neighbor_row = point.row as i32 + shift;

            if west_east_neighbor_col >= 0 && west_east_neighbor_col < col_max as i32 {
                neighbors.push(Point{row: point.row, col: west_east_neighbor_col});
            }

            if north_south_neighbor_row >= 0 && north_south_neighbor_row < row_max as i32 {
                neighbors.push(Point{row: north_south_neighbor_row, col: point.col});
            }
        }

        return neighbors;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn single_island() {
        let grid = vec![
            vec!['1', '1', '1', '1', '0'],
            vec!['1', '1', '0', '1', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '0', '0', '0'],
        ];
        assert_eq!(Solution::num_islands(grid), 1)
    }

    #[test]
    fn all_land() {
        let grid = vec![
            vec!['1', '1', '1', '1', '1'],
            vec!['1', '1', '1', '1', '1'],
            vec!['1', '1', '1', '1', '1'],
            vec!['1', '1', '1', '1', '1'],
        ];
        assert_eq!(Solution::num_islands(grid), 1)
    }

    #[test]
    fn one_island_with_lakes() {
        let grid = vec![
            vec!['1', '1', '1', '1', '1'],
            vec!['1', '0', '1', '0', '1'],
            vec!['1', '1', '1', '1', '1'],
            vec!['1', '0', '1', '0', '1'],
        ];
        assert_eq!(Solution::num_islands(grid), 1)
    }

    #[test]
    fn multiple_islands() {
        let grid = vec![
            vec!['1', '1', '0', '0', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '1', '0', '0'],
            vec!['0', '0', '0', '1', '1'],
        ];
        assert_eq!(Solution::num_islands(grid), 3)
    }

    #[test]
    fn big_grid() {
        let grid = vec![
            vec!['1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','0','1','0','1','1'],
            vec!['0','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','0'],
            vec!['1','0','1','1','1','0','0','1','1','0','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','0','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','0','1','1','1','1','1','1','0','1','1','1','0','1','1','1','0','1','1','1'],
            vec!['0','1','1','1','1','1','1','1','1','1','1','1','0','1','1','0','1','1','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','0','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['0','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','1','1','1','1','0','1','1','1','1','1','1','1','0','1','1','1','1','1','1'],
            vec!['1','0','1','1','1','1','1','0','1','1','1','0','1','1','1','1','0','1','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','0'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','0','0'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'],
            vec!['1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1']
        ];

        assert_eq!(Solution::num_islands(grid), 1)
    }
}

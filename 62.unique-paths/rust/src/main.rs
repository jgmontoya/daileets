fn main() {}

struct Solution {}

impl Solution {
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let mut sol_rows = vec![vec![1; n as usize], vec![0; n as usize]];
        let mut row_to_write = 1;

        for _ in 0..m {
            let row_to_read = 1 - row_to_write;
            for col in 0..n as usize {
                sol_rows[row_to_write][col] = sol_rows[row_to_read][col];
                if col == 0 { continue }

                sol_rows[row_to_write][col] += sol_rows[row_to_write][col - 1];
            }
            row_to_write = row_to_read;
        }

        return sol_rows[row_to_write][(n-1) as usize];
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn small_grid() {
        assert_eq!(Solution::unique_paths(3, 2), 3)
    }

    #[test]
    fn mid_grid() {
        assert_eq!(Solution::unique_paths(3, 7), 28)
    }
}

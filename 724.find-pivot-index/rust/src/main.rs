fn main() {}

struct Solution {}

impl Solution {
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        let mut total = 0;

        for num in &nums {
            total += *num;
        }

        let mut index = 0;
        let mut prefix_sum = 0;
        let mut suffix_sum = total;
        for num in nums {
            suffix_sum -= num;
            if prefix_sum == suffix_sum {
                return index;
            }
            prefix_sum += num;
            index += 1;
        }

        return -1;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn pivot_index_in_the_inside() {
        assert_eq!(Solution::pivot_index(vec![1, 7, 3, 6, 5, 6]), 3);
    }

    #[test]
    fn no_pivot_index() {
        assert_eq!(Solution::pivot_index(vec![1, 2, 3]), -1);
    }

    #[test]
    fn pivot_on_leftmost_index() {
        assert_eq!(Solution::pivot_index(vec![2, 1, -1]), 0);
    }
}

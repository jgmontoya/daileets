struct Solution {}
impl Solution {
    pub fn running_sum(nums: Vec<i32>) -> Vec<i32> {
        let mut solution = Vec::with_capacity(nums.len());

        let mut accum = 0;
        for num in nums {
            accum += num;
            solution.push(accum);
        }
        return solution;
    }
}

fn main() {
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one_to_four() {
        assert_eq!(Solution::running_sum(vec![1, 2, 3, 4]), [1, 3, 6, 10])
    }

    #[test]
    fn five_ones() {
        assert_eq!(Solution::running_sum(vec![1, 1, 1, 1, 1]), [1, 2, 3, 4, 5])
    }

    #[test]
    fn random_ints() {
        assert_eq!(Solution::running_sum(vec![3, 1, 2, 10, 1]), [3, 4, 6, 16, 17])
    }
}

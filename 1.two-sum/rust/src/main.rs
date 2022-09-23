use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut num_to_index: HashMap<i32, usize> = HashMap::new();

        for (index, num) in nums.iter().enumerate() {
            let target_partner = target - num;

            match num_to_index.get(&target_partner) {
                Some(partner_index) => return vec![*partner_index as i32, index as i32],
                None => { num_to_index.insert(*num, index); }
            }
        };

        vec![0, 1]
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn first_example() {
        assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![0, 1])
    }

    #[test]
    fn second_example() {
        assert_eq!(Solution::two_sum(vec![3, 2, 4], 6), vec![1, 2])
    }

    #[test]
    fn third_example() {
        assert_eq!(Solution::two_sum(vec![3, 3], 6), vec![0, 1])
    }
}

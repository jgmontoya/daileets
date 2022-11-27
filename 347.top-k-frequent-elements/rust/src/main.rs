use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut frequency_map = HashMap::with_capacity(nums.len());

        for &num in &nums {
            frequency_map.entry(num).and_modify(|count| *count += 1).or_insert(1);
        }

        let mut capacity_groups: Vec<Vec<i32>> = vec![Vec::with_capacity(nums.len()); nums.len() + 1];

        for (&key, &freq) in frequency_map.iter() {
            capacity_groups[freq as usize].push(key);
        }

        let mut result: Vec<i32> = Vec::with_capacity(k as usize);

        for capacity_group in capacity_groups.iter().rev() {
            if capacity_group.len() == 0 { continue }

            result.extend_from_slice(capacity_group);
            if result.len() == k as usize { return result }
        }

        result
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn with_one_num() {
        assert_eq!(Solution::top_k_frequent(vec![1], 1), vec![1])
    }

    #[test]
    fn with_multiple_nums() {
        assert_eq!(Solution::top_k_frequent(vec![1, 1, 1, 2, 2, 3], 2), vec![1, 2])
    }
}

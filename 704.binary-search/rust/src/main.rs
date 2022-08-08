fn main() {}

struct Solution {}

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        if nums.len() == 0 {
            return -1
        }
        let mid = nums.len() / 2;
        let mid_number = nums[mid];
        if target == mid_number {
            return mid as i32
        }

        if target < mid_number {
            return Solution::search(nums[..mid].to_vec(), target)
        }

        else {
            let subsolution = Solution::search(nums[mid+1..].to_vec(), target);
            if subsolution == -1 {
                return -1;
            }
            return subsolution + mid as i32 + 1
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn return_index_when_num_exists() {
        assert_eq!(Solution::search(vec![-1,0,3,5,9,12], 9), 4)
    }

    #[test]
    fn return_neg_one_when_num_doesnt_exist() {
        assert_eq!(Solution::search(vec![-1,0,3,5,9,12], 2), -1)
    }
}

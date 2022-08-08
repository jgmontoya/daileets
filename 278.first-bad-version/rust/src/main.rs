struct Solution {
    isBadVersion: fn(i32) -> bool,
}

impl Solution {
    fn isBadVersion(version: i32) -> bool {
        return (vec![false, false, false, true, true])[(version as usize) - 1];
    }

    fn recursive_first_bad_version(&self, n: i32, accumulated_version: i32) -> i32 {
        if n == 1 {
            return n + accumulated_version;
        }

        let mid_version = n / 2;

        if (self.isBadVersion)(mid_version + accumulated_version) {
            return self.recursive_first_bad_version(mid_version, accumulated_version)
        } else {
            return self.recursive_first_bad_version(mid_version + 1, accumulated_version + mid_version)
        }
    }

    pub fn first_bad_version(&self, n: i32) -> i32 {
        return self.recursive_first_bad_version(n, 0);
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn returns_first_true_version() {
        let solution = Solution {
            isBadVersion: Solution::isBadVersion,
        };
        assert_eq!(solution.first_bad_version(5), 4);
    }
}

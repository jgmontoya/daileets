use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        let mut letter_count = HashMap::new();

        for &character in s.as_bytes().iter() {
            *letter_count.entry(character).or_insert(0) += 1;
        }

        let mut result = 0;
        let mut odd_included = false;
        for &count in letter_count.values() {
            if count % 2 == 0 {
                result += count;
                continue;
            }

            if !odd_included {
                result += count;
                odd_included = true;
                continue;
            }

            if count > 1 {
                result += count - 1
            }
        }

        return result;
    }
}

fn main() { }

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn multiletter() {
        assert_eq!(Solution::longest_palindrome("abccccdd".to_string()), 7)
    }

    #[test]
    fn single_letter() {
        assert_eq!(Solution::longest_palindrome("a".to_string()), 1)
    }
}

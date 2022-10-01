use std::cmp;
use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn character_replacement(s: String, k: i32) -> i32 {
        let string_bytes = s.as_bytes();
        let mut chars = HashMap::new();
        let mut max_length = 0;
        let mut left_index = 0;
        let mut right_index = 0;

        while right_index < string_bytes.len() {
            let right_index_char = string_bytes[right_index];
            chars.entry(right_index_char).and_modify(|count| *count += 1).or_insert(1);


            let max_freq = Solution::max_freq_in_window(&chars);

            let needed_replacements = (right_index - left_index + 1) as i32 - max_freq;
            if needed_replacements > k {
                let left_index_char = string_bytes[left_index];
                chars.entry(left_index_char).and_modify(|count| *count -= 1);
                left_index += 1;
            }

            max_length = cmp::max(max_length, (right_index - left_index + 1) as i32);
            right_index += 1;
        }

        max_length
    }

    fn max_freq_in_window(chars: &HashMap<u8, i32>) -> i32 {
        let mut max_freq = -1;
        for &freq in chars.values() {
            if freq > max_freq {
                max_freq = freq;
            }
        }

        max_freq
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn first_example() {
        assert_eq!(Solution::character_replacement("ABAB".to_string(), 2), 4)
    }

    #[test]
    fn second_example() {
        assert_eq!(Solution::character_replacement("AABABBA".to_string(), 1), 4)
    }

    #[test]
    fn from_non_initial_char() {
        assert_eq!(Solution::character_replacement("ABABBB".to_string(), 2), 6)
    }

    #[test]
    fn longer_example() {
        let input_string = "KRSCDCSONAJNHLBMDQGIFCPEKPOHQIHLTDIQGEKLRLCQNBOHNDQGHJPNDQPERNFSSSRDE\
        QLFPCCCARFMDLHADJADAGNNSBNCJQOF".to_string();
        assert_eq!(Solution::character_replacement(input_string, 4), 7)
    }
}

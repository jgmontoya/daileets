use std::collections::HashMap;

struct Solution {}
fn main() {}

impl Solution {
    pub fn find_anagrams(s: String, p: String) -> Vec<i32> {
        if p.len() > s.len() { return vec![] }

        let mut needed_letters_count = HashMap::new();
        for char in p.as_bytes() {
            needed_letters_count.entry(char).and_modify(|count| *count += 1).or_insert(1);
        }

        let s_bytes = s.as_bytes();

        let mut left_index = 0;
        let mut right_index = 0;
        let mut result = vec![];
        let mut substring_letter_count = needed_letters_count.clone();

        while right_index < p.len() {
            if Solution::process_right_index(s_bytes, &mut substring_letter_count, right_index) {
                if substring_letter_count.values().all(|&count| count == 0) {
                    result.push(left_index as i32);
                }
            }
            right_index += 1;
        }

        while left_index < s.len() - p.len() {
            Solution::process_right_index(s_bytes, &mut substring_letter_count, right_index);
            right_index += 1;
            Solution::process_left_index(s_bytes, &mut substring_letter_count, left_index);
            left_index += 1;
            if substring_letter_count.values().all(|&count| count == 0) {
                result.push(left_index as i32);
            }
        }


        return result
    }

    fn process_right_index<'a>(s_bytes: &'a [u8], substring_letter_count: &mut HashMap<&'a u8, i32>, right_index: usize) -> bool {
        match substring_letter_count.get(&s_bytes[right_index]) {
            Some(_) => {
                substring_letter_count.entry(&s_bytes[right_index]).and_modify(|count| *count -= 1);
                return true
            }
            None => {}
        }
        return false
    }

    fn process_left_index<'a>(s_bytes: &'a [u8], substring_letter_count: &mut HashMap<&'a u8, i32>, left_index: usize) -> bool {
        match substring_letter_count.get(&s_bytes[left_index]) {
            Some(_) => {
                substring_letter_count.entry(&s_bytes[left_index]).and_modify(|count| *count += 1);
                return true
            }
            None => {}
        }
        return false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn first_example() {
        assert_eq!(Solution::find_anagrams("cbaebabacd".to_string(), "abc".to_string()), vec![0, 6])
    }

    #[test]
    fn second_example() {
        assert_eq!(Solution::find_anagrams("abab".to_string(), "ab".to_string()), vec![0, 1, 2])
    }

    #[test]
    fn third_example() {
        assert_eq!(Solution::find_anagrams("baa".to_string(), "aa".to_string()), vec![1])
    }

    #[test]
    fn fourth_example() {
        assert_eq!(Solution::find_anagrams("abacbabc".to_string(), "abc".to_string()), vec![1, 2, 3, 5])
    }
}

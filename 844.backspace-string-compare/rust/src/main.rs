struct Solution {}

impl Solution {
    pub fn backspace_compare(s: String, t: String) -> bool {
        let s_bytes = s.as_bytes();
        let t_bytes = t.as_bytes();

        let mut s_index = (s_bytes.len() - 1) as i32;
        let mut t_index = (t_bytes.len() - 1) as i32;

        let backspace_char: u8 = *"#".to_string().as_bytes().first().unwrap();

        let mut s_skip = 0;
        let mut t_skip = 0;

        while s_index >= 0 || t_index >= 0 {
            let s_char;
            let t_char;

            if s_index >= 0 {
                s_char = s_bytes[s_index as usize];
            } else {
                s_char = 0;
            }

            if s_char == backspace_char {
                s_index -= 1;
                s_skip += 1;
                continue;
            }

            if s_skip > 0 {
                s_index -= 1;
                s_skip -= 1;
                continue
            }

            if t_index >= 0 {
                t_char = t_bytes[t_index as usize];
            } else {
                t_char = 0;
            }

            if t_char == backspace_char {
                t_index -= 1;
                t_skip += 1;
                continue;
            }

            if t_skip > 0 {
                t_index -= 1;
                t_skip -= 1;
                continue;
            }

            if s_char != t_char { return false }
            s_index -= 1;
            t_index -= 1;
        }

        return true
    }

    // pub fn stack_backspace_compare(s: String, t: String) -> bool {
    //     let mut max_stack = Vec::with_capacity(200);
    //     let mut min_stack = Vec::with_capacity(200);

    //     let max_bytes;
    //     let min_bytes;

    //     if s.len() >= t.len() {
    //         max_bytes = s.as_bytes();
    //         min_bytes = t.as_bytes();
    //     } else {
    //         max_bytes = t.as_bytes();
    //         min_bytes = s.as_bytes();
    //     }

    //     let backspace_char = *"#".to_string().as_bytes().first().unwrap();

    //     for (index, &max_char) in max_bytes.iter().enumerate() {
    //         if max_char != backspace_char {
    //             max_stack.push(max_char);
    //         } else {
    //             max_stack.pop();
    //         }

    //         if index >= min_bytes.len() { continue }

    //         let min_char = min_bytes[index];
    //         if min_char != backspace_char {
    //             min_stack.push(min_char);
    //         } else {
    //             min_stack.pop();
    //         }
    //     }

    //     return max_stack == min_stack
    // }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn simple_backspace_true() {
        assert!(Solution::backspace_compare("ab#c".to_string(), "ad#c".to_string()))
    }

    #[test]
    fn empty_result_true() {
        assert!(Solution::backspace_compare("ab##".to_string(), "c#d#".to_string()))
    }

    #[test]
    fn simple_backspace_false() {
        assert_eq!(Solution::backspace_compare("a#c".to_string(), "b".to_string()), false)
    }
}

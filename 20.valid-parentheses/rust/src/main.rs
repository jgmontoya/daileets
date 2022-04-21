struct Solution {
}

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let input_length = s.len();
        if input_length % 2 == 1 {
            return false
        }

        let mut stack = Vec::with_capacity(input_length/2);

        for c in s.chars() {
            if c == '(' || c == '[' || c == '{' {
                stack.push(c);
                if stack.len() > input_length / 2 {
                    return false
                }
            } else {
                let top = stack.pop();
                let top_char = match top {
                    Some(character) => character,
                    None => return false
                };
                if Solution::parenthesis_mismatch(top_char, c) {
                    return false
                }
            }
        }

        stack.len() == 0
    }

    fn parenthesis_mismatch(p1: char, p2: char) -> bool {
        if p1 == '(' {
            return p2 != ')'
        } else if p1 == '[' {
            return p2 != ']'
        } else {
            return p2 != '}'
        }
    }
}

fn main() {
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn true_on_good_single_pair() {
        assert_eq!(Solution::is_valid("()".to_string()), true);
    }

    #[test]
    fn true_on_multiple_good_pairs() {
        assert_eq!(Solution::is_valid("()[]{}".to_string()), true);
    }

    #[test]
    fn false_on_unmatching_pair() {
        assert_eq!(Solution::is_valid("(]".to_string()), false);
    }

    #[test]
    fn false_on_odd_length_input() {
        assert_eq!(Solution::is_valid("(".to_string()), false)
    }

    #[test]
    fn false_on_only_opening() {
        assert_eq!(Solution::is_valid("((".to_string()), false)
    }
}

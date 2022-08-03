struct Solution {}

impl Solution {
    pub fn is_subsequence(s: String, t: String) -> bool {
        return Solution::_is_subsequence(s.as_bytes(), t.as_bytes());
    }

    fn _is_subsequence(s: &[u8], t: &[u8]) -> bool {
        if s.len() > t.len() {
            return false;
        }

        if s.len() == 0 {
            return true
        }

        if s[0] == t[0] {
            return Solution::_is_subsequence(&s[1..], &t[1..])
        }

        return Solution::_is_subsequence(&s, &t[1..])
    }
}

fn main() { }

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn is_subsequence() {
        assert_eq!(Solution::is_subsequence("abc".to_string(), "ahbgdc".to_string()), true)
    }

    #[test]
    fn not_subsequence() {
        assert_eq!(Solution::is_subsequence("axc".to_string(), "ahbgdc".to_string()), false)
    }
}

use std::collections::HashMap;

struct Solution { }

impl Solution {
    pub fn is_isomorphic(s: String, t: String) -> bool {
        let mut isomorphism = HashMap::new();
        let mut mapped_from = HashMap::new();

        for (index, character) in s.chars().enumerate() {
            let other_char = t.chars().nth(index).unwrap();

            match isomorphism.get(&character) {
                Some(isomorphic_char) => {
                    if isomorphic_char != &other_char {
                        return false;
                    } else {
                        if !Solution::check_many_to_one(&mapped_from, other_char, character) {
                            return false;
                        }
                    }
                },
                None => {
                    if !Solution::check_many_to_one(&mapped_from, other_char, character) {
                        return false;
                    }
                    isomorphism.insert(character, other_char);
                    mapped_from.insert(other_char, character);
                }
            }
        }

        return true
    }

    fn check_many_to_one(mapped_from: &HashMap<char, char>, other_char: char, character: char) -> bool {
        match mapped_from.get(&other_char) {
            Some(original_char) => {
                if &character != original_char {
                    return false;
                }
            },
            None => { }
        }
        return true;
    }
}


fn main() { }

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn simple_isomorphism() {
        let egg = "egg".to_string();
        let add = "add".to_string();
        assert_eq!(Solution::is_isomorphic(egg, add), true);
    }

    #[test]
    fn no_isomorphism_one_to_many() {
        let foo = "foo".to_string();
        let bar = "bar".to_string();
        assert_eq!(Solution::is_isomorphic(foo, bar), false);
    }

    #[test]
    fn longer_isomorphism() {
        let paper = "paper".to_string();
        let title = "title".to_string();
        assert_eq!(Solution::is_isomorphic(paper, title), true);
    }

    #[test]
    fn no_isomorphism_many_to_one() {
        let abc = "abc".to_string();
        let aba = "aba".to_string();
        assert_eq!(Solution::is_isomorphic(abc, aba), false);
    }
}

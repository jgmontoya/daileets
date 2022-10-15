use std::{collections::HashMap, cmp};

struct Solution {}

impl Solution {
    pub fn get_hint(secret: String, guess: String) -> String {
        let mut bulls = 0;
        let mut cows = 0;

        let secret_bytes = secret.as_bytes();
        let guess_bytes = guess.as_bytes();

        let mut digits = HashMap::new();

        for (index, &secret_digit) in secret_bytes.iter().enumerate() {
            let guess_digit = guess_bytes[index];
            if guess_digit == secret_digit {
                bulls += 1;
                continue;
            }

            digits.entry(secret_digit).and_modify(|counter: &mut (i32, i32)| (*counter).0 += 1).or_insert((1, 0));
            digits.entry(guess_digit).and_modify(|counter| (*counter).1 += 1).or_insert((0, 1));
        }

        for (secret_count, guess_count) in digits.values() {
            cows += cmp::min(secret_count, guess_count)
        }

        return format!("{bulls}A{cows}B")
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one_bull_three_cows() {
        assert_eq!(Solution::get_hint("1807".to_string(), "7810".to_string()), "1A3B".to_string())
    }

    #[test]
    fn one_bull_one_cow() {
        assert_eq!(Solution::get_hint("1123".to_string(), "0111".to_string()), "1A1B".to_string())
    }
}

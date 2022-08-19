fn main() {}

struct Solution {}

impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        let mut first = 0;
        let mut second = 1;
        let mut temp: i32;
        for _ in 0..n {
            temp = first + second;
            first = second;
            second = temp;
        }
        return second;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn two_steps() {
        assert_eq!(Solution::climb_stairs(2), 2)
    }

    #[test]
    fn three_steps() {
        assert_eq!(Solution::climb_stairs(3), 3)
    }
}

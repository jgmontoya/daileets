fn main() {}

struct Solution {}

impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n < 2 { return n }

        let mut first = 0;
        let mut second = 1;
        let mut temp: i32;

        for _ in 0..(n-1) {
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
    fn first_number() {
        assert_eq!(Solution::fib(0), 0)
    }

    #[test]
    fn second_number() {
        assert_eq!(Solution::fib(1), 1)
    }

    #[test]
    fn third_number() {
        assert_eq!(Solution::fib(2), 1)
    }

    #[test]
    fn fourth_number() {
        assert_eq!(Solution::fib(3), 2)
    }

    #[test]
    fn fifth_number() {
        assert_eq!(Solution::fib(4), 3)
    }

    #[test]
    fn large_number() {
        assert_eq!(Solution::fib(20), 6765)
    }


}

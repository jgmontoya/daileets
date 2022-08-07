use std::cmp;
struct Solution {}

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        if prices.len() < 2 {
            return 0
        }

        if prices.len() == 2 {
            return cmp::max(0, prices[1] - prices[0])
        }

        let mut max_profit: i32;
        let mut selected_buy_price: i32;
        let mut selected_sell_price: i32;

        if prices[0] <= prices[1] {
            selected_buy_price = prices[0];
            selected_sell_price = prices[1];
        } else {
            selected_buy_price = prices[1];
            selected_sell_price = prices[2];
        }

        max_profit = selected_sell_price - selected_buy_price;
        for (_index, &value) in prices[2..].iter().enumerate() {
            let index = _index + 2;
            if value > selected_sell_price {
                selected_sell_price = value;
            } else if value < selected_buy_price {
                selected_buy_price = value;
                if index < prices.len() - 1 {
                    selected_sell_price = prices[index + 1];
                } else {
                    selected_sell_price = -1;
                }
            }
            max_profit = cmp::max(max_profit, selected_sell_price - selected_buy_price);
        }

        return cmp::max(0, max_profit);
    }
}
fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn some_profit() {
        assert_eq!(Solution::max_profit(vec![7, 1, 5, 3, 6, 4]), 5);
    }

    #[test]
    fn buy_day_0_profit() {
        assert_eq!(Solution::max_profit(vec![1,2,4,2,5,7,2,4,9,0]
        ), 8);
    }

    #[test]
    fn no_profit() {
        assert_eq!(Solution::max_profit(vec![7, 6, 4, 3, 1]), 0);
    }
}

use std::cmp;

fn main() {}
struct Solution {}

impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let mut min_costs = vec![0; cost.len()];

        min_costs[0] = cost[0];
        min_costs[1] = cost[1];

        for (step, step_cost) in cost.iter().enumerate() {
            if step < 2 { continue }

            let single_step = min_costs[step-1] + step_cost;
            let double_step = min_costs[step-2] + step_cost;
            min_costs[step] = std::cmp::min(single_step, double_step);
        }

        return std::cmp::min(min_costs[cost.len()-1], min_costs[cost.len()-2]);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn simple_start_at_index1() {
        assert_eq!(Solution::min_cost_climbing_stairs(vec![10, 15, 20]), 15)
    }

    #[test]
    fn multiple_hops_start_at_index0() {
        assert_eq!(Solution::min_cost_climbing_stairs(vec![1, 100, 1, 1, 1, 100, 1, 1, 100, 1]), 6)
    }

    #[test]
    fn just_two_steps() {
        assert_eq!(Solution::min_cost_climbing_stairs(vec![1, 2]), 1)
    }
}

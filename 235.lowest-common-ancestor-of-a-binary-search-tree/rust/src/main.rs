use std::rc::Rc;
use std::cell::RefCell;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
  pub val: i32,
  pub left: Option<Rc<RefCell<TreeNode>>>,
  pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    TreeNode {
      val,
      left: None,
      right: None
    }
  }
}
struct Solution {}

impl Solution {
    pub fn lowest_common_ancestor(
        root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if p == root || q == root {
            return root
        }

        let root_node = root.clone().unwrap();
        let root_val = root_node.borrow().val;

        let p_node = p.clone().unwrap();
        let p_val = p_node.borrow().val;
        let q_node = q.clone().unwrap();
        let q_val = q_node.borrow().val;

        if (p_val < root_val && q_val > root_val) || (q_val < root_val && p_val > root_val) {
            return root;
        }

        if p_val < root_val {
            let left_son = root.clone().unwrap().borrow().left.to_owned();
            return Solution::lowest_common_ancestor(left_son, p, q);
        }

        let right_son = root.clone().unwrap().borrow().right.to_owned();
        return Solution::lowest_common_ancestor(right_son, p, q);
    }
}
fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn simple_root_solution() {
        let son = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: None, right: None})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 2, left: son.clone(), right: None})));
        assert_eq!(Solution::lowest_common_ancestor(root.clone(), root.clone(), son), root);
    }
}

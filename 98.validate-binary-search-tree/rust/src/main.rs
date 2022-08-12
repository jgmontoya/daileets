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
    pub fn is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        Solution::recursive_is_valid_bst(root, None, None)
    }

    fn recursive_is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>, min: Option<i32>, max: Option<i32>) -> bool {
        if root.is_none() {
            return true
        }

        let root_node = root.unwrap();
        let root_val = root_node.borrow().val;

        if let Some(min_val) = min {
            if root_val <= min_val {
                return false
            }
        }

        if let Some(max_val) = max {
            if root_val >= max_val {
                return false
            }
        }

        if let Some(left_son) = &root_node.borrow().left {
            if left_son.borrow().val >= root_val {
                return false
            }
            if !Solution::recursive_is_valid_bst(root_node.borrow().left.clone(), min, Some(root_val)) {
                return false
            }
        }

        if let Some(right_son) = &root_node.borrow().right {
            if right_son.borrow().val <= root_val {
                return false
            }
            if !Solution::recursive_is_valid_bst(root_node.borrow().right.clone(), Some(root_val), max) {
                return false
            }
        }
        return true
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn valid_tree() {
        let left = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: None, right: None})));
        let right = Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 2, left: left, right: right})));
        assert_eq!(Solution::is_valid_bst(root), true);
    }

    #[test]
    fn direct_child_violation_tree() {
        let left_grandson = Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None})));
        let right_grandson = Some(Rc::new(RefCell::new(TreeNode{val: 6, left: None, right: None})));
        let left_son = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: None, right: None})));
        let right_son = Some(Rc::new(RefCell::new(TreeNode{val: 4, left: left_grandson, right: right_grandson})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 5, left: left_son, right: right_son})));
        assert_eq!(Solution::is_valid_bst(root), false);
    }

    #[test]
    fn ancester_violation_tree() {
        let left_grandson = Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None})));
        let right_grandson = Some(Rc::new(RefCell::new(TreeNode{val: 7, left: None, right: None})));
        let left_son = Some(Rc::new(RefCell::new(TreeNode{val: 4, left: None, right: None})));
        let right_son = Some(Rc::new(RefCell::new(TreeNode{val: 6, left: left_grandson, right: right_grandson})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 5, left: left_son, right: right_son})));
        assert_eq!(Solution::is_valid_bst(root), false);
    }
}

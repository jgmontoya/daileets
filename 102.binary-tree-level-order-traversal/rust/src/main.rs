
use std::rc::Rc;
use std::cell::RefCell;
use std::collections::VecDeque;

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
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        if root.is_none() {
            return Vec::new();
        }

        let mut result: Vec<Vec<i32>> = Vec::new();
        let mut node_queue = VecDeque::new();
        node_queue.push_back((0, root));

        while node_queue.len() > 0 {
            if let Some((level, Some(node))) = node_queue.pop_front() {
                if level >= result.len() {
                    result.push(vec![]);
                }
                result[level].push(node.borrow().val);
                let left_child = node.borrow().left.clone();
                node_queue.push_back((level + 1, left_child));
                let right_child = node.borrow().right.clone();
                node_queue.push_back((level + 1, right_child));
            }
        }

        return result;
    }
}

fn main() {}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn empty_array_for_null_tree() {
        let result: Vec<Vec<i32>> = Vec::new();
        assert_eq!(Solution::level_order(None), result);
    }

    #[test]
    fn val_array_for_single_node() {
        let node = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: None, right: None})));
        assert_eq!(Solution::level_order(node), vec![vec![1]]);
    }

    #[test]
    fn proper_vec_for_tree() {
        let grandson = Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None})));
        let son = Some(Rc::new(RefCell::new(TreeNode{val: 2, left: grandson, right: None})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: None, right: son})));
        assert_eq!(Solution::level_order(root), vec![vec![1], vec![2], vec![3]]);
    }

    #[test]
    fn proper_vec_for_left_line_tree() {
        let grand3son = Some(Rc::new(RefCell::new(TreeNode{val: 5, left: None, right: None})));
        let grandgrandson = Some(Rc::new(RefCell::new(TreeNode{val: 4, left: grand3son, right: None})));
        let grandson = Some(Rc::new(RefCell::new(TreeNode{val: 3, left: grandgrandson, right: None})));
        let son = Some(Rc::new(RefCell::new(TreeNode{val: 2, left: grandson, right: None})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: son, right: None})));
        assert_eq!(Solution::level_order(root), vec![vec![1], vec![2], vec![3], vec![4], vec![5]]);
    }

    #[test]
    fn multibranch_tree() {
        let grandson1 = Some(Rc::new(RefCell::new(TreeNode{val: 15, left: None, right: None})));
        let grandson2 = Some(Rc::new(RefCell::new(TreeNode{val: 7, left: None, right: None})));
        let son1 = Some(Rc::new(RefCell::new(TreeNode{val: 9, left: None, right: None})));
        let son2 = Some(Rc::new(RefCell::new(TreeNode{val: 20, left: grandson1, right: grandson2})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 3, left: son1, right: son2})));
        assert_eq!(Solution::level_order(root), vec![vec![3], vec![9,20], vec![15,7]]);
    }
}

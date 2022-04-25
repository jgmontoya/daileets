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
struct Solution {
}
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result = &mut Vec::with_capacity(100);
        Solution::recursive_inorder_traversal(&root, result);
        return result.to_vec()
    }

    fn recursive_inorder_traversal(root: &Option<Rc<RefCell<TreeNode>>>, result: &mut Vec<i32>) {
        if let Some(node) = root {
            let node = node.borrow();
            Solution::recursive_inorder_traversal(&node.left, result);
            result.push(node.val);
            Solution::recursive_inorder_traversal(&node.right, result);
        }
    }
}

fn main() {
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn empty_array_for_null_tree() {
        assert_eq!(Solution::inorder_traversal(None), []);
    }

    #[test]
    fn val_number_for_single_node() {
        let node = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: None, right: None})));
        assert_eq!(Solution::inorder_traversal(node), vec![1]);
    }

    #[test]
    fn proper_vec_for_tree() {
        let grandson = Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None})));
        let son = Some(Rc::new(RefCell::new(TreeNode{val: 2, left: grandson, right: None})));
        let root = Some(Rc::new(RefCell::new(TreeNode{val: 1, left: None, right: son})));
        assert_eq!(Solution::inorder_traversal(root), vec![1, 3, 2]);
    }
}

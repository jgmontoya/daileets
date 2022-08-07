#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

struct Solution {}

impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }

        let mut current: Option<Box<ListNode>> = None;
        let mut center = head.as_ref();

        while !center.is_none() {
            if let Some(ref center_node) = center {
                let mut new_node = ListNode::new(center_node.val);

                new_node.next = current;

                current = Some(Box::new(new_node));
                center = center_node.next.as_ref();
            }
        }
        return current;
    }
}
fn main() { }

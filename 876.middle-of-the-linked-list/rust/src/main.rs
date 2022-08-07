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

struct Solution { }

impl Solution {
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut fast = head.as_ref();
        let mut slow = head.clone();

        let mut counter = 0;
        while !fast.is_none() {
            fast = fast.unwrap().next.as_ref();
            if counter % 2 != 0 {
                slow = slow.unwrap().next;
            }
            counter += 1;
        }

        return slow;
    }
}

fn main() { }

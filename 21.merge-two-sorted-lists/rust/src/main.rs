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
  pub fn merge_two_lists(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if list1.is_none() {
      return list2;
    }
    if list2.is_none() {
      return list1;
    }

    let (
      head_val,
      list1,
      list2
    ) = Solution::pick_min_and_advance(list1, list2);

    return Some(Box::new(ListNode {
      val: head_val,
      next: Solution::merge_two_lists(list1, list2)
    }));
  }

  fn pick_min_and_advance(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> (
    i32, Option<Box<ListNode>>, Option<Box<ListNode>>
  ) {
    if list1.as_ref().unwrap().val <= list2.as_ref().unwrap().val {
      return (list1.as_ref().unwrap().val, list1.unwrap().next, list2)
    }

    (list2.as_ref().unwrap().val, list1, list2.unwrap().next)
  }
}

fn main() { }

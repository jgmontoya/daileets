
class ListNode {
  val: number
  next: ListNode | null

  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}

function detectCycle(head: ListNode | null): ListNode | null {
  if (head === null) return null

  let slow: ListNode | null = head;
  let fast = head;

  let new_fast = fast?.next?.next;
  if (new_fast === undefined || new_fast == null) return null;

  fast = new_fast;
  slow = slow.next as ListNode;

  while (slow !== fast) {
    let new_fast = fast?.next?.next;
    if (new_fast === undefined || new_fast == null) return null;

    fast = new_fast;
    slow = slow.next as ListNode;
  }

  while (head !== slow) {
    slow = slow.next as ListNode;
    head = head.next as ListNode;
  }

  return slow;
};

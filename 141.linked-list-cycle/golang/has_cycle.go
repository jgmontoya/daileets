package has_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	for {
		if fast.Next == nil || slow.Next == nil {
			return false
		}

		if fast.Next.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
}

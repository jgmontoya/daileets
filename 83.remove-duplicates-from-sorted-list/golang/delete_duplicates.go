package delete_duplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	for head != nil && head.Next != nil && head.Next.Val == head.Val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	head.Next = deleteDuplicates(head.Next)
	return head
}

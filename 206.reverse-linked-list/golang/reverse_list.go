package reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var left *ListNode
	center := head
	right := center.Next

	for right != nil {
		center.Next = left
		left = center
		center = right
		right = right.Next
	}
	center.Next = left

	return center
}

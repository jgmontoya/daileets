package middle_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	counter := 0
	for fast != nil {
		fast = fast.Next
		if counter%2 != 0 {
			slow = slow.Next
		}
		counter++
	}
	return slow
}

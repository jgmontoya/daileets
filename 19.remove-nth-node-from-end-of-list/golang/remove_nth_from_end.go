package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	ahead_node := head
	node_to_remove := head

	for counter := 1; counter <= n; counter++ {
		ahead_node = ahead_node.Next
	}

	node_to_remove_started := false
	previous_to_remove := node_to_remove

	for ahead_node != nil {
		ahead_node = ahead_node.Next
		node_to_remove = node_to_remove.Next
		if node_to_remove_started {
			previous_to_remove = previous_to_remove.Next
		}
		node_to_remove_started = true
	}

	if !node_to_remove_started {
		return head.Next
	}

	previous_to_remove.Next = node_to_remove.Next
	return head
}

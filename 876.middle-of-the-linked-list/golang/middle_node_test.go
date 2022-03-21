package middle_node

import "testing"

func ListNodeFromIntArray(int_array []int) *ListNode {
	var last_node *ListNode
	for index := len(int_array) - 1; index >= 0; index-- {
		last_node = &ListNode{Val: int_array[index], Next: last_node}
	}
	return last_node
}

func CheckMiddleNode(t *testing.T, input *ListNode, solution *ListNode) {
	middle_node := MiddleNode(input)
	if middle_node == solution {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", middle_node.Val, solution.Val)
		t.Fail()
	}
}

func TestMiddleNode(t *testing.T) {
	input := ListNodeFromIntArray([]int{1, 2, 3, 4, 5})
	solution := input.Next.Next
	CheckMiddleNode(t, input, solution)

	input = ListNodeFromIntArray([]int{1, 2, 3, 4, 5, 6})
	solution = input.Next.Next.Next
	CheckMiddleNode(t, input, solution)
}

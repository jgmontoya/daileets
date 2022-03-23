package remove_nth_from_end

import "testing"

func SliceEqual(first_slice []int, second_slice []int) bool {
	first_length := len(first_slice)
	second_length := len(second_slice)

	if first_length != second_length {
		return false
	}

	for i := 0; i < first_length; i++ {
		if first_slice[i] != second_slice[i] {
			return false
		}
	}
	return true
}

func IntArrayFromListNode(head *ListNode) []int {
	var array []int
	for node := head; node != nil; node = node.Next {
		array = append(array, node.Val)
	}
	return array
}

func ListNodeFromIntArray(int_array []int) *ListNode {
	var last_node *ListNode
	for index := len(int_array) - 1; index >= 0; index-- {
		last_node = &ListNode{Val: int_array[index], Next: last_node}
	}
	return last_node
}

func CheckMiddleNode(t *testing.T, input *ListNode, n int, solution []int) {
	removed_list := RemoveNthFromEnd(input, n)
	removed_list_array := IntArrayFromListNode(removed_list)
	if SliceEqual(removed_list_array, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", removed_list_array, solution)
		t.Fail()
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	input_list := ListNodeFromIntArray([]int{1, 2, 3, 4, 5})
	n := 2
	solution := []int{1, 2, 3, 5}
	CheckMiddleNode(t, input_list, n, solution)

	input_list = ListNodeFromIntArray([]int{1})
	n = 1
	solution = []int{}
	CheckMiddleNode(t, input_list, n, solution)

	input_list = ListNodeFromIntArray([]int{1, 2})
	n = 1
	solution = []int{1}
	CheckMiddleNode(t, input_list, n, solution)

	input_list = ListNodeFromIntArray([]int{1, 2})
	n = 2
	solution = []int{2}
	CheckMiddleNode(t, input_list, n, solution)
}

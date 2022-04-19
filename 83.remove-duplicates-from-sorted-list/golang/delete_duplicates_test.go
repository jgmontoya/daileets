package delete_duplicates

import "testing"

func array2List(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}

	nodes := []*ListNode{}
	for _, val := range array {
		node := &ListNode{Val: val}
		nodes = append(nodes, node)
	}
	for index, node := range nodes {
		if index == len(nodes)-1 {
			break
		}
		node.Next = nodes[index+1]
	}
	return nodes[0]
}

func list2Array(head *ListNode) []int {
	array := []int{}
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	return array
}

func LinkedListEqual(list1 *ListNode, list2 *ListNode) bool {
	if list1 == nil {
		return list2 == nil
	}
	if list2 == nil {
		return false
	}
	if list1.Val != list2.Val {
		return false
	}
	return LinkedListEqual(list1.Next, list2.Next)
}

func CheckDeleteDuplicates(t *testing.T, input []int, solution []int) {
	deleted_duplicates := deleteDuplicates(array2List(input))
	if LinkedListEqual(deleted_duplicates, array2List(solution)) {
		t.Log("OK")
	} else {
		t.Errorf("deleteDuplicates(%d): Got %d; expected %d", input, list2Array(deleted_duplicates), solution)
		t.Fail()
	}
}

func TestDeleteDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	solution := []int{1, 2}
	CheckDeleteDuplicates(t, input, solution)

	input = []int{1, 1, 2, 3, 3}
	solution = []int{1, 2, 3}
	CheckDeleteDuplicates(t, input, solution)
}

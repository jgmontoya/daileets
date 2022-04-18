package remove_elements

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

func CheckRemoveElements(t *testing.T, input_list []int, val int, solution []int) {
	remove_elements := removeElements(array2List(input_list), val)
	if LinkedListEqual(remove_elements, array2List(solution)) {
		t.Log("OK")
	} else {
		t.Errorf("remove_elements(%d, %d): Got %d; expected %d", input_list, val, list2Array(remove_elements), solution)
		t.Fail()
	}
}

func TestRemoveElements(t *testing.T) {
	input_list := []int{1, 2, 6, 3, 4, 5, 6}
	val := 6
	solution := []int{1, 2, 3, 4, 5}
	CheckRemoveElements(t, input_list, val, solution)

	input_list = []int{}
	val = 1
	solution = []int{}
	CheckRemoveElements(t, input_list, val, solution)

	input_list = []int{7, 7, 7, 7}
	val = 7
	solution = []int{}
	CheckRemoveElements(t, input_list, val, solution)
}

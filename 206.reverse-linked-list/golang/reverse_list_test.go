package reverse_list

import "testing"

func listFromIntSlice(int_array []int) *ListNode {
	var list_node *ListNode
	for index := len(int_array) - 1; index >= 0; index-- {
		list_node = &ListNode{Val: int_array[index], Next: list_node}
	}
	return list_node
}

func listEquals(list1 *ListNode, list2 *ListNode) bool {
	if list1 == nil {
		return list2 == nil
	} else if list2 == nil {
		return false
	}

	if list1.Val != list2.Val {
		return false
	}

	return listEquals(list1.Next, list2.Next)
}

func intSliceFromList(list *ListNode) []int {
	slice := []int{}
	for list != nil {
		slice = append(slice, list.Val)
		list = list.Next
	}
	return slice
}

func checkReverseList(t *testing.T, input *ListNode, solution *ListNode) {
	reversed_list := reverseList(input)
	if listEquals(reversed_list, solution) {
		t.Log("OK")
	} else {
		t.Errorf("Got %d; expected %d", intSliceFromList(reversed_list), intSliceFromList(solution))
		t.Fail()
	}
}
func TestReverseList(t *testing.T) {
	input := listFromIntSlice([]int{1, 2, 3, 4, 5})
	solution := listFromIntSlice([]int{5, 4, 3, 2, 1})
	checkReverseList(t, input, solution)

	input = listFromIntSlice([]int{1, 2})
	solution = listFromIntSlice([]int{2, 1})
	checkReverseList(t, input, solution)

	input = listFromIntSlice([]int{})
	solution = listFromIntSlice([]int{})
	checkReverseList(t, input, solution)
}

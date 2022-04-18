package has_cycle

import "testing"

func linkedListFromArray(array []int, pos int) *ListNode {
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
			if pos != -1 {
				node.Next = nodes[pos]
			}
			continue
		}
		node.Next = nodes[index+1]
	}
	return nodes[0]
}

func CheckHasCycle(t *testing.T, head *ListNode, solution bool) {
	has_cycle := hasCycle(head)
	if has_cycle == solution {
		t.Log("OK")
	} else {
		t.Errorf("hasCycle(%d): Got %t; expected %t", head.Val, has_cycle, solution)
	}
}

func TestHasCycle(t *testing.T) {
	list, pos := []int{3, 2, 0, -4}, 1
	CheckHasCycle(t, linkedListFromArray(list, pos), pos != -1)

	list, pos = []int{1, 2}, 0
	CheckHasCycle(t, linkedListFromArray(list, pos), pos != -1)

	list, pos = []int{1}, -1
	CheckHasCycle(t, linkedListFromArray(list, pos), pos != -1)

	list, pos = []int{1, 2}, -1
	CheckHasCycle(t, linkedListFromArray(list, pos), pos != -1)

	list, pos = []int{}, -1
	CheckHasCycle(t, linkedListFromArray(list, pos), pos != -1)
}

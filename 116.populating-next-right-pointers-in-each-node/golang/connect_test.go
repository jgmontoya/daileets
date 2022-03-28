package connect

import (
	"strconv"
	"testing"
)

func TreeFromSlice(slice []int) *Node {
	if len(slice) == 0 {
		return nil
	}
	root := Node{}
	root.Val = slice[0]

	AddSons(&root, slice, 0)
	return &root
}

func AddSons(node *Node, slice []int, index int) {
	left_index := 2*index + 1
	right_index := 2*index + 2
	if left_index < len(slice) && slice[left_index] != -1 {
		(*node).Left = &Node{Val: slice[left_index]}
		AddSons((*node).Left, slice, left_index)
	}
	if right_index < len(slice) && slice[right_index] != -1 {
		(*node).Right = &Node{Val: slice[right_index]}
		AddSons((*node).Right, slice, right_index)
	}
}

func SliceFromTree(root *Node) []string {
	if root == nil {
		return []string{}
	}

	result := []string{}
	node_queue := []*Node{}
	node_queue = append(node_queue, root)

	for len(node_queue) != 0 {
		node := Pop(&node_queue)
		if node.Val == -1 {
			continue
		}
		result = append(result, strconv.Itoa(node.Val))
		if node.Left != nil {
			node_queue = append(node_queue, node.Left)
		} else {
			node_queue = append(node_queue, &Node{Val: -1})
		}

		if node.Right != nil {
			node_queue = append(node_queue, node.Right)
		} else {
			node_queue = append(node_queue, &Node{Val: -1})
		}

		if node.Next == nil {
			result = append(result, "#")
		}
	}
	return result
}

func SliceEqual(first_slice []string, second_slice []string) bool {
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

func CheckConnect(t *testing.T, input *Node, solution []string) {
	connect := SliceFromTree(Connect(input))
	if SliceEqual(connect, solution) {
		t.Logf("OK")
	} else {
		t.Errorf("Got %s; expected %s", connect, solution)
		t.Fail()
	}
}
func TestConnect(t *testing.T) {
	input := TreeFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	solution := []string{"1", "#", "2", "3", "#", "4", "5", "6", "7", "#"}
	CheckConnect(t, input, solution)

	input = TreeFromSlice([]int{})
	solution = []string{}
	CheckConnect(t, input, solution)
}

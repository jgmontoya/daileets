package invert_tree

import "testing"

func AddSons(node *TreeNode, slice []int, index int) {
	left_index := 2*index + 1
	right_index := 2*index + 2
	if left_index < len(slice) && slice[left_index] != -1 {
		(*node).Left = &TreeNode{Val: slice[left_index]}
		AddSons((*node).Left, slice, left_index)
	}
	if right_index < len(slice) && slice[right_index] != -1 {
		(*node).Right = &TreeNode{Val: slice[right_index]}
		AddSons((*node).Right, slice, right_index)
	}
}

func TreeFromSlice(slice []int) *TreeNode {
	if len(slice) == 0 {
		return nil
	}
	root := TreeNode{}
	root.Val = slice[0]

	AddSons(&root, slice, 0)
	return &root
}

func EqualTrees(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil {
		if root2 == nil {
			return true
		}
		return false
	}
	if root1.Val != root2.Val || (root1.Left != nil && root2.Left == nil) || (root1.Right != nil && root2.Right == nil) {
		return false
	}

	if root1.Left != nil {
		if !EqualTrees(root1.Left, root2.Left) {
			return false
		}
	}

	if root1.Right != nil {
		if !EqualTrees(root1.Right, root2.Right) {
			return false
		}
	}
	return true
}

func Pop(queue *[]*TreeNode) *TreeNode {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	return node
}

func SliceFromTree(root *TreeNode) []int {
	result := []int{}
	node_queue := []*TreeNode{}
	node_queue = append(node_queue, root)

	for len(node_queue) != 0 {
		node := Pop(&node_queue)
		if node.Val == -1 {
			continue
		}
		result = append(result, node.Val)
		if node.Left != nil {
			node_queue = append(node_queue, node.Left)
		} else {
			node_queue = append(node_queue, &TreeNode{Val: -1})
		}

		if node.Right != nil {
			node_queue = append(node_queue, node.Right)
		} else {
			node_queue = append(node_queue, &TreeNode{Val: -1})
		}
	}
	return result
}

func CheckInvertTree(t *testing.T, input *TreeNode, solution *TreeNode) {
	inverted_tree := InvertTree(input)
	if EqualTrees(inverted_tree, solution) {
		t.Logf("OK")
	} else {
		t.Errorf("Got %d; expected %d", SliceFromTree(inverted_tree), SliceFromTree(solution))
		t.Fail()
	}
}

func TestInvertTree(t *testing.T) {
	input := TreeFromSlice([]int{4, 2, 7, 1, 3, 6, 9})
	solution := TreeFromSlice([]int{4, 7, 2, 9, 6, 3, 1})
	CheckInvertTree(t, input, solution)

	input = TreeFromSlice([]int{2, 1, 3})
	solution = TreeFromSlice([]int{2, 3, 1})
	CheckInvertTree(t, input, solution)

	input = TreeFromSlice([]int{})
	solution = TreeFromSlice([]int{})
	CheckInvertTree(t, input, solution)
}

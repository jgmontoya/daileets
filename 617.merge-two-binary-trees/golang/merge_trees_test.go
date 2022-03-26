package merge_trees

import (
	"testing"
)

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

func CheckMergeTrees(t *testing.T, root1 *TreeNode, root2 *TreeNode, solution *TreeNode) {
	merged_trees := MergeTrees(root1, root2)
	if EqualTrees(merged_trees, solution) {
		t.Logf("OK")
	} else {
		t.Errorf("Got %d; expected %d", SliceFromTree(merged_trees), SliceFromTree(solution))
		t.Fail()
	}
}

func TestMergeTrees(t *testing.T) {
	root1 := TreeFromSlice([]int{1, 3, 2, 5})
	root2 := TreeFromSlice([]int{2, 1, 3, -1, 4, -1, 7})
	solution := TreeFromSlice([]int{3, 4, 5, 5, 4, -1, 7})
	CheckMergeTrees(t, root1, root2, solution)

	root1 = TreeFromSlice([]int{1})
	root2 = TreeFromSlice([]int{1, 2})
	solution = TreeFromSlice([]int{2, 2})
	CheckMergeTrees(t, root1, root2, solution)

	root1 = TreeFromSlice([]int{})
	root2 = TreeFromSlice([]int{1})
	solution = TreeFromSlice([]int{1})
	CheckMergeTrees(t, root1, root2, solution)
}

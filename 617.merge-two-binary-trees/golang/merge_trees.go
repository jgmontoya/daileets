package merge_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	merged_trees := &TreeNode{Val: root1.Val + root2.Val}
	if root1.Left != nil {
		if root2.Left != nil {
			merged_trees.Left = MergeTrees(root1.Left, root2.Left)
		} else {
			merged_trees.Left = root1.Left
		}
	} else if root2.Left != nil {
		merged_trees.Left = root2.Left
	}

	if root1.Right != nil {
		if root2.Right != nil {
			merged_trees.Right = MergeTrees(root1.Right, root2.Right)
		} else {
			merged_trees.Right = root1.Right
		}
	} else if root2.Right != nil {
		merged_trees.Right = root2.Right
	}

	return merged_trees
}

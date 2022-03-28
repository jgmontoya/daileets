package connect

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return root
	}
	node_queue := []*Node{}
	node_queue = append(node_queue, root)
	node_count := 0
	next_level := 1

	for len(node_queue) != 0 {
		node := Pop(&node_queue)
		node_count++

		if node.Left != nil {
			node_queue = append(node_queue, node.Left)
			node_queue = append(node_queue, node.Right)
		}

		if node_count == next_level {
			next_level = 2*next_level + 1
		} else {
			node.Next = node_queue[0]
		}

	}
	return root
}

func Pop(queue *[]*Node) *Node {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	return node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var queue []*Node
	queue = append(queue, root)

	for len(queue) > 0 {
		lvlSize := len(queue)

		for i := range lvlSize {
			node := queue[0]
			queue = queue[1:]
			if i != lvlSize-1 {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	stack := make([]*TreeNode, 0)

	if root != nil {
		stack = append(stack, root.Left)
		stack = append(stack, root.Right)
		for len(stack) > 0 {
			node1 := stack[len(stack)-2]
			node2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			switch {
			case node1 == nil && node2 == nil:
				continue
			case node1 == nil || node2 == nil:
				return false
			default:
				if node1.Val != node2.Val {
					return false
				}
			}
			stack = append(stack, node1.Left)
			stack = append(stack, node2.Right)
			stack = append(stack, node1.Right)
			stack = append(stack, node2.Left)
		}
	}

	return true
}
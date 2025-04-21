/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {

	var d int

	var recFuncDepth func(root *TreeNode) int

	recFuncDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := recFuncDepth(node.Left)
		rightDepth := recFuncDepth(node.Right)
		currD := leftDepth + rightDepth
		d = max(d, currD)
		return max(leftDepth, rightDepth) + 1
	}

	recFuncDepth(root)

	return d
}
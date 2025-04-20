/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	var depthFunc func(*TreeNode) (status bool, depth int)

	depthFunc = func(node *TreeNode) (status bool, depth int) {
		if node == nil {
			return true, 0
		}
		statL, depthL := depthFunc(node.Left)
		statR, depthR := depthFunc(node.Right)
		return statL && statR && depthL-depthR <= 1 && depthL-depthR >= -1, max(depthL, depthR) + 1
	}
	statRoot, _ := depthFunc(root)
	return statRoot
}
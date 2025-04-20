/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int
	var recFunc func(node *TreeNode, lvl int)
	recFunc = func(node *TreeNode, lvl int) {
		if node == nil {
			return
		}
		if lvl > len(res)-1 {
			res = append(res, node.Val)
		}
		recFunc(node.Right, lvl+1)
		recFunc(node.Left, lvl+1)
	}
	recFunc(root, 0)
	return res
}
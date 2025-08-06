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

	var seacher func(node *TreeNode, lvl int)

	seacher = func(node *TreeNode, lvl int) {
		if node == nil {
			return
		}
		if len(res) < lvl {
			res = append(res, node.Val)
		}
		seacher(node.Right, lvl+1)
		seacher(node.Left, lvl+1)
	}
	seacher(root, 1)
	return res
}
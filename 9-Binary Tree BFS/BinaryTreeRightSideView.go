/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		res = append(res, queue[0].Val)
		lvlSize := len(queue)
		for _ = range lvlSize {
			node := queue[0]
			queue = queue[1:]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}
	return res
}
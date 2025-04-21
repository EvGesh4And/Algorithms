/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	var res int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		lvlSize := len(queue)
		res = 0
		for _ = range lvlSize {
			node := queue[0]
			queue = queue[1:]
			res += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
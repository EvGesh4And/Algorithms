/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var res []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		lvlSize := len(queue)
		maxValLevel := queue[0].Val
		for _ = range lvlSize {
			node := queue[0]
			queue = queue[1:]
			if maxValLevel < node.Val {
				maxValLevel = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, maxValLevel)
	}
	return res
}
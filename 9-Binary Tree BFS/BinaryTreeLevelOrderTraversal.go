/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var res [][]int
	queue = append(queue, root)
	lvl := 0
	for len(queue) > 0 {
		lvlsize := len(queue)
		res = append(res, []int{})
		for _ = range lvlsize {
			node := queue[0]
			queue = queue[1:]
			res[lvl] = append(res[lvl], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		lvl++
	}
	return res
}
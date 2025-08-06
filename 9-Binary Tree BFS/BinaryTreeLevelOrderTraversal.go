/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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

	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		lvlSize := len(queue)
		lvlSlice := make([]int, 0, lvlSize)
		for _ = range lvlSize {
			node := queue[0]
			queue = queue[1:]
			lvlSlice = append(lvlSlice, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, lvlSlice)
	}
	return res
}
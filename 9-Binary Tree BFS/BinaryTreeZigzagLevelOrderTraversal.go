/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int

	var queue []*TreeNode
	queue = append(queue, root)

	flag := true
	for len(queue) > 0 {
		lvlSize := len(queue)
		lvlSlice := make([]int, lvlSize)
		for i := range lvlSize {
			node := queue[0]
			queue = queue[1:]
			if flag {
				lvlSlice[i] = node.Val
			} else {
				lvlSlice[lvlSize-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, lvlSlice)
		flag = !flag
	}
	return res
}
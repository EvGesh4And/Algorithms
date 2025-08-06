/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	type st struct {
		node   *TreeNode
		target int
	}

	stack := make([]st, 0)
	stack = append(stack, st{root, targetSum})

	for len(stack) > 0 {
		stCurr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if stCurr.node == nil {
			continue
		}
		currTarget := stCurr.target - stCurr.node.Val
		if stCurr.node.Left == nil && stCurr.node.Right == nil {
			if currTarget == 0 {
				return true
			}
			continue
		}
		stack = append(stack, st{stCurr.node.Left, currTarget})
		stack = append(stack, st{stCurr.node.Right, currTarget})
	}
	return false
}
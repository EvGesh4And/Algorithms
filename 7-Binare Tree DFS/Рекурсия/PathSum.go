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
	currSum := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return currSum == 0
	}
	return hasPathSum(root.Left, currSum) || hasPathSum(root.Right, currSum)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	type nodeDepth struct {
		node  *TreeNode
		depth int
	}

	stack := make([]nodeDepth, 0)

	maxDepth := 0

	stack = append(stack, nodeDepth{root, 0})

	for len(stack) > 0 {
		st := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if st.node != nil {
			stack = append(stack, nodeDepth{st.node.Left, st.depth + 1}, nodeDepth{st.node.Right, st.depth + 1})
		} else {
			maxDepth = max(maxDepth, st.depth)
		}
	}
	return maxDepth
}
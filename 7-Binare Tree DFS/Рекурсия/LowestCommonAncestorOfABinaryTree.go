/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var dfs func(node, p, q *TreeNode) (bool, *TreeNode)

	dfs = func(node, p, q *TreeNode) (bool, *TreeNode) {
		if node == nil {
			return false, nil
		}

		currStatus := node == p || node == q

		leftStatus, leftLCA := dfs(node.Left, p, q)
		rightStatus, rightLCA := dfs(node.Right, p, q)

		// Если LCA уже найден в левом или правом поддереве
		if leftLCA != nil {
			return true, leftLCA
		}
		if rightLCA != nil {
			return true, rightLCA
		}
		// Если текущий узел — LCA
		if leftStatus && rightStatus || ((leftStatus || rightStatus) && currStatus) {
			return true, node
		}
		// Иначе, возвращаем, найден ли кто-то из p/q
		return currStatus || leftStatus || rightStatus, nil
	}

	_, lca := dfs(root, p, q)

	return lca
}
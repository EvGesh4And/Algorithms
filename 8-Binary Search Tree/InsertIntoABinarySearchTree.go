/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	currNode := root

	for {
		if currNode.Val == val {
			break
		}
		if val < currNode.Val {
			if currNode.Left != nil {
				currNode = currNode.Left
			} else {
				currNode.Left = &TreeNode{Val: val}
				break
			}
		} else {
			if currNode.Right != nil {
				currNode = currNode.Right
			} else {
				currNode.Right = &TreeNode{Val: val}
				break
			}
		}

	}

	return root
}
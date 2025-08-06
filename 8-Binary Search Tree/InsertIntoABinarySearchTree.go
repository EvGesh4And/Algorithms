/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Это решение красивее
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

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
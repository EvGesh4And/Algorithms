/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

// Лучшее решение
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	return helper(node.Left, min, &node.Val) && helper(node.Right, &node.Val, max)
}

type elem struct {
	node      *TreeNode
	limitDown int
	limitUp   int
}

func isValidBST(root *TreeNode) bool {
	var stack []elem
	stack = append(stack, elem{root, math.MinInt64, math.MaxInt64})
	for len(stack) > 0 {
		currElem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if currElem.node != nil {
			if currElem.node.Val <= currElem.limitDown || currElem.node.Val >= currElem.limitUp {
				return false
			}
			stack = append(stack, elem{currElem.node.Left, currElem.limitDown, currElem.node.Val})
			stack = append(stack, elem{currElem.node.Right, currElem.node.Val, currElem.limitUp})
		}
	}
	return true
}
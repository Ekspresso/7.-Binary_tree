// Функция вычисляет сумму значений
// самых глубоких узлов бинарного дерева.

package binarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	return sumToMax(root, maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		l := maxDepth(root.Left) + 1
		r := maxDepth(root.Right) + 1
		if l > r {
			return l
		} else {
			return r
		}
	}
	return 0
}

func sumToMax(root *TreeNode, max int) int {
	if root != nil {
		left := sumToMax(root.Left, max-1)
		right := sumToMax(root.Right, max-1)
		if max == 1 {
			return root.Val
		}
		return left + right
	}
	return 0
}

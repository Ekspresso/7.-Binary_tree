// Функция вычисляет максимальную глубину бинарного дерева.

package binarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

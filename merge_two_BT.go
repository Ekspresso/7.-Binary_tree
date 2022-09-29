// Функция сливает 2 бинарных дерева:
// если в обеих деревьях есть этот узел, то она складывает их значения;
// если ветка есть в одном дереве, а в другом её нет,
// то дерево просто наращивается этой веткой.
// Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// Output: [3,4,5,5,4,null,7]

package binarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция видоизменяет дерево 1.
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if (root1 == nil && root2 == nil) || (root1 != nil && root2 == nil) {
		return root1
	} else if root1 == nil && root2 != nil {
		return root2
	}
	root1.Val += root2.Val
	mergeTrees(root1.Left, root2.Left)
	mergeTrees(root1.Right, root2.Right)
	if root1.Left == nil && root2.Left != nil {
		root1.Left = root2.Left
	}
	if root1.Right == nil && root2.Right != nil {
		root1.Right = root2.Right
	}
	return root1
}

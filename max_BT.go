// Функция создаёт из массива максимальное бинарное дерево:
// 1. находит максимальный элемент,
// 2. все элементы массива слева от макс значения,
// должны находиться слева в дереве, а все элементы справа,
// соответственно справа в дереве.
// 3. Каждый родительский узел в дереве больше всех дочерних.

package binarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) > 0 {
		max := findMax(nums)
		root := new(TreeNode)
		root.Val = nums[max]
		root.Left = constructMaximumBinaryTree(nums[:max])
		root.Right = constructMaximumBinaryTree(nums[max+1:])
		return root
	}
	return nil
}

func findMax(nums []int) int {
	max := -1
	ind := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			ind = i
		}
	}
	return ind
}

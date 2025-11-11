package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func reverseInorderTraversal(root *BinaryTree) []int32 {
	result := make([]int32, 0)
	if root != nil {
		result = append(result, reverseInorderTraversal(root.right)...)
		result = append(result, root.data.value)
		result = append(result, reverseInorderTraversal(root.left)...)
	}
	return result
}

// 逆間順
// 右→自分→左

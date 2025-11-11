package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func inorderTraversal(root *BinaryTree) []int32 {
	result := make([]int32, 0)
	if root != nil {
		result = append(result, inorderTraversal(root.left)...)
		result = append(result, root.data.value)
		result = append(result, inorderTraversal(root.right)...)
	}

	return result
}

// 間順走査
// = 昇順になるやつ

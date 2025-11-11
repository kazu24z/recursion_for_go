package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func postorderTraversal(root *BinaryTree) []int32 {
	result := make([]int32, 0)
	if root != nil {
		result = append(result, postorderTraversal(root.left)...)
		result = append(result, postorderTraversal(root.right)...)
		result = append(result, root.data.value)
	}
	return result
}

// 後順走査
// = 削除しやすいやつ
// 左→右→自分

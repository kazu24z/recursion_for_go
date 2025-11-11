package main

// import (
//     "fmt"
//     "strings"
// )

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func preorderTraversal(root *BinaryTree) []int32 {
	result := make([]int32, 0)
	if root != nil {
		result = append(result, root.data.value)
		result = append(result, preorderTraversal(root.left)...)
		result = append(result, preorderTraversal(root.right)...)
	}

	return result
}

// 前順走査
// → 自分→左→右

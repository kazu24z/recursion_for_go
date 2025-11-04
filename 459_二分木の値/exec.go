package main

// import (
//     "fmt"
// )

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func sumOfThreeNodes(root *BinaryTree) int32 {
	if root == nil {
		return 0
	}

	rootData := root.data
	leftValue := int32(0)
	if root.left != nil {
		leftValue = root.left.data.value
	}

	rightValue := int32(0)
	if root.right != nil {
		rightValue = root.right.data.value
	}

	return rootData.value + leftValue + rightValue
}

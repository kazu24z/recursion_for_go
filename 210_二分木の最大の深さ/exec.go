package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func maximumDepth(root *BinaryTree) int32 {

	if root == nil {
		return -1
	}

	counterLeft := maximumDepth(root.left)
	counterRight := maximumDepth(root.right)

	if counterLeft >= counterRight {
		return counterLeft + 1
	} else {
		return counterRight + 1
	}
}

// 二分木の最大の深さ = 最長経路
// 前順でtreeを再現していく流れの中で、root == nilになるまでをサイクルとみてカウントしていく？？？

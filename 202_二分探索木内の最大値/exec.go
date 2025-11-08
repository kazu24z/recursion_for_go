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

func maximumNode(root *BinaryTree) *BinaryTree {
	if root == nil {
		return nil
	}

	iterator := root
	for iterator.right != nil {
		iterator = iterator.right
	}

	return iterator
}

// 平衡二分探索木の最大値 = 一番右側 → ひたすら右に

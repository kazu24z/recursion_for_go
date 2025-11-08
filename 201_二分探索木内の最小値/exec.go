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

func minimumNode(root *BinaryTree) *BinaryTree {
	if root == nil {
		return nil
	}

	iterator := root
	for iterator.left != nil {
		iterator = iterator.left
	}
	return iterator
}

// 平衡二分探索木において最小値 = 一番左ノード(leaf)
// ひたすらleftを掘っていって、leftがnilになったらOK

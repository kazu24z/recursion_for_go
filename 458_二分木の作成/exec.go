package main

import "fmt"

func main() {
	binaryTree := &BinaryTree{data: 10}

	node2 := &BinaryTree{data: 6}
	node3 := &BinaryTree{data: 8}

	binaryTree.left = node2
	binaryTree.right = node3

	fmt.Println(binaryTree.data)
	fmt.Println(binaryTree.left.data)
	fmt.Println(binaryTree.right.data)
}

type BinaryTree struct {
	data  int32
	left  *BinaryTree
	right *BinaryTree
}

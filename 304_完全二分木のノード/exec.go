package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func countNodes(root *BinaryTree) int32 {
	if root == nil {
		return 0
	}

	l := countNodes(root.left)
	r := countNodes(root.right)

	return l + r + 1
}

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

func exists(root *BinaryTree, key int32) bool {

	for n := root; n != nil; {
		if key == n.data.value {
			return true
		}

		if key < n.data.value {
			n = n.left
		} else {
			n = n.right
		}
	}

	return false
}

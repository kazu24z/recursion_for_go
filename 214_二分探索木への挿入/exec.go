package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func bstInsert(root *BinaryTree, key int32) *BinaryTree {
	iterator := root

	for iterator != nil {
		if key == iterator.data.value {
			return root
		}
		if key < iterator.data.value {
			if iterator.left == nil {
				iterator.left = &BinaryTree{data: &Integer{value: key}}
				break
			}
			iterator = iterator.left
		} else {
			if iterator.right == nil {
				iterator.right = &BinaryTree{data: &Integer{value: key}}
				break
			}
			iterator = iterator.right
		}
	}
	return root
}

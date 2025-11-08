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

func bstSearch(root *BinaryTree, key int32) *BinaryTree {
	iterator := root
	for iterator != nil {
		// 条件合致したらその時点でreturn
		if iterator.data.value == key {
			return iterator
		}
		// もしroot.data.value > keyならleft
		if iterator.data.value >= key {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	// 最後まで探索したくせに見つからないならnil
	return nil
}

// アホ！イテレーションと再帰を混ぜて使ってる!!!
// func bstSearch(root *BinaryTree, key int32) *BinaryTree {
// 	// 関数を完成させてください
// 	if root.data.value == key {
// 		return root
// 	}

// 	iterator := root
// 	for iterator != nil {
// 		// もしroot.data.value > keyならleft
// 		if iterator.data.value >= key {
// 			iterator = iterator.left
// 			bstSearch(iterator, key)
// 		} else {
// 			iterator = iterator.right
// 			bstSearch(iterator, key)
// 		}
// 	}

// 	return nil
// }

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

func successor(root *BinaryTree, key int32) *BinaryTree {
	if root == nil {
		return nil
	}

	node := search(root, key)
	// keyがヒットしなかった場合
	if node == nil {
		return nil
	}

	ite := node
	// ノードにrightがある場合は、1回右→後ひたすら左
	if node.right != nil {
		ite = ite.right
		// 後ひたすら左
		for ite.left != nil {
			ite = ite.left
		}
	} else {
		// ite.rightがnillの場合
		// 最後にleftしたときのnodeが対象
		ite = searchLastLeftNode(root, key)

	}

	return ite
}

// 前提として、
func searchLastLeftNode(root *BinaryTree, key int32) *BinaryTree {
	// leftしたときのノードを保存する箱（上書き）
	var leftDownNode *BinaryTree

	iterator := root
	for iterator != nil {
		if key == iterator.data.value {
			return leftDownNode
		}

		if key < iterator.data.value {
			leftDownNode = iterator
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return leftDownNode
}

// 復習がてsearchを実装する
func search(root *BinaryTree, key int32) *BinaryTree {
	// イテレータを用意
	iterator := root

	// 各回で、keyがその回のdataより小さいならleftへ、大きいならrightへ
	for iterator != nil {
		if key == iterator.data.value {
			return iterator
		}

		if key <= iterator.data.value {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return nil
}

// 右にノードがある場合

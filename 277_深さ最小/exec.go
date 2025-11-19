package main

import "fmt"

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func main() {
	tree := toBinaryTree([]int32{97, 10, 77, 32, 40, 70, 32, 96, 27, 10, 12, 93, 82, 33, 55, 71, 59, 82, 37, 75, 25, 31, 14, 96, 85, 41, 28, 70, 9, 56, 8, 90, 8, 65, 49, 45, 34, 30, 25, 7, 7, 97, 23, 66, 84, 57, 38, 38, 95, 9, 77, 60, 44, 3, 81, 41, 89, 90, 73, 100, 86, 53, 96, 40})
	min := minDepth(tree)

	fmt.Println(min)
}

func minDepth(root *BinaryTree) int32 {
	level := minDepthHelper(root, 0)

	return level
}

func minDepthHelper(root *BinaryTree, level int32) int32 {
	// ベースケース
	if root == nil {
		return 0
	}
	// 「葉」のときは
	if root.left == nil && root.right == nil {
		return 0
	}

	leftLevel := minDepthHelper(root.left, level)
	rightLevel := minDepthHelper(root.right, level)

	// 子が片方だけの場合
	if root.left == nil {
		return rightLevel + 1
	}

	if root.right == nil {
		return leftLevel + 1
	}

	// 子が左右存在の場合
	if leftLevel < rightLevel {
		return leftLevel + 1
	} else {
		return rightLevel + 1
	}
}

// 再帰処理で、部分木単位で左右比較して階層が少ないほうが勝ち を繰り返す
func toBinaryTree(arr []int32) *BinaryTree {
	if len(arr) == 0 {
		return nil
	}
	return arrayToBinaryTreeHelper(arr, 0)
}

func arrayToBinaryTreeHelper(arr []int32, index int) *BinaryTree {
	if index >= len(arr) {
		return nil
	}

	node := &BinaryTree{
		data:  &Integer{value: arr[index]},
		left:  arrayToBinaryTreeHelper(arr, 2*index+1), // 左子
		right: arrayToBinaryTreeHelper(arr, 2*index+2), // 右子
	}

	return node
}

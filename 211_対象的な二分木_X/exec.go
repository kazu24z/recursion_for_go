package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func symmetricTree(root *BinaryTree) bool {
	if root == nil {
		return true
	}

	return helper(root.left, root.right)

}

func helper(left, right *BinaryTree) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.data.value == right.data.value &&
		helper(left.left, right.right) && // 左の左 ↔ 右の右
		helper(left.right, right.left)

}

// 各階層において、自分のright, leftのノード = 同じ階層のノードのleft ,rightが同じ

package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func deepestLeaves(root *BinaryTree) int32 {
	// ベースケース
	if root == nil {
		return 0
	}
	sum, _ := deepestLeavesHelper(root, 0)
	return sum
}

func deepestLeavesHelper(root *BinaryTree, level int32) (int32, int32) {
	// ベースケース
	if root == nil {
		return 0, -1
	}

	// もし自分のleft, rightがnilなら = 葉 →
	if root.left == nil && root.right == nil {
		return root.data.value, level
	}

	left, leftLevel := deepestLeavesHelper(root.left, level+1)
	right, rightLevel := deepestLeavesHelper(root.right, level+1)
	if leftLevel > rightLevel {
		return left, leftLevel
	} else if leftLevel < rightLevel {
		return right, rightLevel
	} else {
		return left + right, leftLevel
	}
}

// 葉:left, rightがともにnilのもの
// 深層部の葉: 一番深い階層の葉のみ

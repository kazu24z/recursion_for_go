package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func treeWithTheSameValue(root *BinaryTree) bool {
	if root == nil {
		return true
	}

	result := treeWithTheSameValueHelper(root, root.data.value)

	return result
}

// ヘルパ関数
func treeWithTheSameValueHelper(root *BinaryTree, value int32) bool {

	if root == nil {
		return true
	}

	// これ、自分自身はどうなん？？って判定
	if root.data.value != value {
		return false
	}

	// ここは「子」階層の結果やねん
	leftResult := treeWithTheSameValueHelper(root.left, value)
	rightResult := treeWithTheSameValueHelper(root.right, value)

	// これは、上で自分自身はチェック済みだけど、子もOKじゃないとだめなので子をチェックしてる
	return leftResult && rightResult
}

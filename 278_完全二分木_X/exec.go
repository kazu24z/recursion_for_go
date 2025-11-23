package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func isCompleteBinaryTree(root *BinaryTree) bool {
	nodeCount := countNodes(root)
	return isCompleteHelper(root, 0, nodeCount)
}

func countNodes(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.left) + countNodes(root.right)
}

func isCompleteHelper(root *BinaryTree, index int, nodeCount int) bool {
	if root == nil {
		return true
	}

	// インデックスがノード数以上になったら完全二分木ではない
	if index >= nodeCount {
		return false
	}

	// 左右の子を再帰的にチェック
	return isCompleteHelper(root.left, 2*index+1, nodeCount) &&
		isCompleteHelper(root.right, 2*index+2, nodeCount)
}

/**

NG: 非Leafで左右どちらか欠けている
NG: Leafで、左がないのに右だけある
→ 一番下の葉とそれ以外で処理をわける

Leafの条件：right, leftともにnil

*/

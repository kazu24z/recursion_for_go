package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func validateBST(root *BinaryTree) bool {
	return ValidateBSTHelper(root, nil, nil)
}

// func validateBSTHelper(root *BinaryTree, arr *[]int32) {
// 	if root == nil {
// 		return
// 	}

// 	validateBSTHelper(root.left, arr)
// 	*arr = append(*arr, root.data.value)
// 	validateBSTHelper(root.right, arr)
// }

func ValidateBSTHelper(root, min, max *BinaryTree) bool {
	if root == nil {
		return true
	}

	// 左の場合
	if min != nil && root.data.value <= min.data.value {
		return false
	}
	// 右の場合
	if max != nil && root.data.value >= max.data.value {
		return false
	}

	return ValidateBSTHelper(root.left, min, root) &&
		ValidateBSTHelper(root.right, root, max)
}

/**
ツリーを入力して、それがBSTですか？って検証する処理
→ 各部分木がさ、BSTであればよい?
→ これはNG。全体を見る必要がある

BSTの条件：ルートからの並びが、親より小さいなら左側、大きいなら右側にならんでますよ
→ ルートから左右に再帰的に

*/

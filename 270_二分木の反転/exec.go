package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func invertTree(root *BinaryTree) *BinaryTree {
	// とりあえず前順走査を実装する
	if root == nil {
		return nil
	}

	left := invertTree(root.left)
	right := invertTree(root.right)

	root.left = right
	root.right = left

	return root
}

// 反転：左右のノードを反転させるってこと
/**
考え方的には、前順で自分→左→右で走査して、左ときは右に、右のときは左にノードを入れる？？
→条件がcurrentNode.value < targetValue なら左（通常の逆）

考え方的には、各部分木において左右を反転させるだけ

*/

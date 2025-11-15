package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func mergeBST(root1 *BinaryTree, root2 *BinaryTree) *BinaryTree {
	if root1 == nil && root2 == nil {
		return nil
	}

	// 片方がnilの場合の処理 => 素直にもう片方を返せばいいだけ
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	// 両方ともnilでない場合
	left := mergeBST(root1.left, root2.left)
	right := mergeBST(root1.right, root2.right)
	sumValue := root1.data.value + root2.data.value

	return &BinaryTree{data: &Integer{value: sumValue}, left: left, right: right}
}

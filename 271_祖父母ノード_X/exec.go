package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func totalEvenGrandparent(root *BinaryTree) int32 {
	// 関数を完成させてください
	sum := totalEvenGrandparentHelper(root, nil, nil)

	return sum
}

func totalEvenGrandparentHelper(root, parent, grandparent *BinaryTree) int32 {
	if root == nil {
		return 0
	}

	left := totalEvenGrandparentHelper(root.left, root, parent)
	right := totalEvenGrandparentHelper(root.right, root, parent)

	value := int32(0)
	if grandparent != nil && grandparent.data.value%2 == 0 {
		value = root.data.value
	}

	return left + right + value
}

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

// 親の親が偶数であるノードの合計値

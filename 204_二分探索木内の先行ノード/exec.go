package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func predecessor(root *BinaryTree, key int32) *BinaryTree {
	node := search(root, key)

	var result *BinaryTree
	if node.left != nil {
		result = maximumNode(node.left)
	} else {
		// keyノードに左側がない場合
		// 1つ上（親）が、さらにその上のノードから見て右だったらこたえ
		iterator := root
		for iterator != nil {
			if iterator.data.value == key {
				return result
			}
			// 左に曲がる
			if key < iterator.data.value {
				iterator = iterator.left
			} else {
				// 右に曲がる
				result = iterator
				iterator = iterator.right
			}
		}
		return nil
	}

	return result
}

func search(root *BinaryTree, key int32) *BinaryTree {
	if root == nil {
		return nil
	}
	iterator := root
	for iterator != nil {
		if iterator.data.value == key {
			return iterator
		}

		if key < iterator.data.value {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return nil
}

func maximumNode(root *BinaryTree) *BinaryTree {
	if root == nil {
		return nil
	}

	iterator := root
	for iterator.right != nil {
		iterator = iterator.right
	}

	return iterator
}

// keyを探索してそれを見つけたあと、その左側にノードがあればそれ以下の部分木の最大値が答え
// keyノードの左側にノードがなかった場合、１つ上（親）から見て、keyノードが右側にあれば答えなければ、さらに上に遡り、現在ノードから見て子が右に来たらその親が答え

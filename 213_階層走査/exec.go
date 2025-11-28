package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func levelOrderTraversal(root *BinaryTree) []*int32 {
	if root == nil {
		return []*int32{}
	}
	deck := make([]*BinaryTree, 0)
	deck = append(deck, root) // rootだけ入れておく
	result := make([]*int32, 0)

	for len(deck) > 0 {
		// queueから取り出す
		poped := deck[0]
		deck = deck[1:]
		// 取り出したノードの値を結果配列に入れる
		if poped == nil {
			result = append(result, nil)
		} else {
			result = append(result, &poped.data.value)

			// 次のノードをqueueに入れる
			deck = append(deck, poped.left)
			deck = append(deck, poped.right)
		}
	}

	return result
}

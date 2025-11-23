package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func rightLevelNode(root *BinaryTree) []int32 {
	maxLevel := int32(0)
	rightestMap := make(map[int32]int32)
	resultArr := make([]int32, 0)
	rightLevelNodeHelper(root, 0, rightestMap, &maxLevel)

	// mapを配列に変換する
	for i := int32(0); i <= maxLevel; i++ {
		resultArr = append(resultArr, rightestMap[i])
	}

	return resultArr

}

func rightLevelNodeHelper(root *BinaryTree, level int32, rightestMap map[int32]int32, maxLevel *int32) {
	if root == nil {
		return
	}

	if level > *maxLevel {
		*maxLevel = level
	}

	_, ok := rightestMap[level]
	if !ok {
		rightestMap[level] = root.data.value
	}

	rightLevelNodeHelper(root.right, level+1, rightestMap, maxLevel)
	rightLevelNodeHelper(root.left, level+1, rightestMap, maxLevel)

}

// 各階層の右側... = rightを代入したほうが答えの候補？？

/**
深さ優先探索は同じ階層の他ノードについては知りようがない（兄弟ノードのみ）
→ って考えると、各階層ごとのノード情報を他で持っておく必要があるよな...
→ 例えば配列に左から入れていって、末尾が一番右って感じで
→ mapかな？
sample[1]: {Node1, Node2, Node3}
sample[2]: {Node5, Node7, Node9}
→ 全部はいらない、各階層のmapにおいて既にノードがあれば入れない。なければ入れる
*/

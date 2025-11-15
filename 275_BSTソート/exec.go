package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func allElementsSorted(root1 *BinaryTree, root2 *BinaryTree) []int32 {

	// 空のスライス
	var arr1 []int32
	var arr2 []int32

	toInOrderArr(root1, &arr1)
	toInOrderArr(root2, &arr2)

	// 2つのスライス同士で分割統治法（マージソート）

	combined := make([]int32, 0)
	i, j := 0, 0
	for len(arr1) > i && len(arr2) > j {
		if arr1[i] < arr2[j] {
			combined = append(combined, arr1[i])
			i++
		} else {
			combined = append(combined, arr2[j])
			j++
		}
	}

	combined = append(combined, arr1[i:]...)
	combined = append(combined, arr2[j:]...)

	return combined
}

func toInOrderArr(root *BinaryTree, arr *[]int32) {
	if root == nil {
		return
	}

	toInOrderArr(root.left, arr)
	*arr = append(*arr, root.data.value)
	toInOrderArr(root.right, arr)

}

// 考え方：1, 2 ともに間順走査（左→自分→右）で昇順にする
// → 両方ともがソート済み配列なるので、マージソートで勝ち抜け処理

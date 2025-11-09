package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func sortedArrToBST(numberList []int32) *BinaryTree {
	return sortedArrToBSTHelper(numberList, 0, int32(len(numberList)-1))
}

func sortedArrToBSTHelper(numberList []int32, start int32, end int32) *BinaryTree {
	// ベースケース
	if start > end {
		return nil
	}
	// 中央mid
	mid := start + (end-start)/2 // start + endにしてないのはint32のオーバフローを防ぐため

	left := sortedArrToBSTHelper(numberList, start, mid-1)
	right := sortedArrToBSTHelper(numberList, mid+1, end)

	root := &BinaryTree{data: &Integer{value: numberList[mid]}, left: left, right: right}
	return root
}

// これ、ソートされてなかったらマージソートとかでソートしてから↑の実装をする
// マージソート: 分割統治法で、配列を半分して言って、分解できなくなったところからleft, rightで比較して勝ち抜け方式
// で、結合配列に入れていく（スタックからPop）。で、最後あまりをappendして完了。計算量O(n * log n)

package main

func topKElements(intArr []int32, k int32) []int32 {
	result := heapsortDesc(intArr)
	return result[:k]
}

// ヒープソートだと昇順になる
// [1, 2, 4, 5, 19]

func heapsortDesc(intArr []int32) []int32 {
	// 入力配列をまず最小ヒープ化
	maxHeap := buildMinHeap(intArr)

	// ヒープのサイズ
	heapSize := int32(len(maxHeap))
	// ===== 繰り返し処理 =======
	return helper(maxHeap, heapSize)

}

func helper(intArr []int32, heapSize int32) []int32 {
	if heapSize <= 1 {
		return intArr
	}
	// 末尾と先頭を入れ替え
	intArr[heapSize-1], intArr[0] = intArr[0], intArr[heapSize-1]

	// 処理範囲を更新
	heapSize--

	// ルートから最大ヒープ化（範囲を限定）
	minHeapify(intArr, 0, heapSize)

	// 再帰的に実行
	return helper(intArr, heapSize)
}

// 配列全体を最小ヒープ化
func buildMinHeap(arr []int32) []int32 {
	// もと配列を変えない
	result := make([]int32, len(arr))
	copy(result, arr)

	mid := int32((len(result) - 1) / 2)

	for i := mid; i >= 0; i-- {
		minHeapify(result, i, int32(len(result)))
	}
	return result
}

func minHeapify(arr []int32, i int32, heapSize int32) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < heapSize && arr[left] < arr[largest] {
		largest = left
	}

	if right < heapSize && arr[right] < arr[largest] {
		largest = right
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		minHeapify(arr, largest, heapSize)
	}
}

/**
1. ルートと末尾を入れ替える
2. 末尾を範囲外にする
3. ルートから全体を最大ヒープする（範囲外を除く）
4. 1~3を繰り返す

*/

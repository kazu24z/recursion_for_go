package main

func minKElements(intArr []int32, k int32) []int32 {
	result := heapsort(intArr)
	return result[:k]
}

func heapsort(intArr []int32) []int32 {
	// 入力配列をまず最小ヒープ化
	maxHeap := buildMaxHeap(intArr)

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
	maxHeapify(intArr, 0, heapSize)

	// 再帰的に実行
	return helper(intArr, heapSize)
}

// 配列全体を最大ヒープ化
func buildMaxHeap(arr []int32) []int32 {
	// もと配列を変えない
	result := make([]int32, len(arr))
	copy(result, arr)

	mid := int32(len(arr)-1) / 2

	for i := mid; i >= 0; i-- {
		maxHeapify(result, i, int32(len(result)))
	}

	return result
}

// 部分木で最大ヒープ
func maxHeapify(arr []int32, i int32, heapSize int32) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}

	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest, heapSize)
	}
}

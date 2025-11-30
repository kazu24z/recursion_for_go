package main

func buildMaxHeap(intArr []int32) []int32 {
	//もとの配列をコピー
	target := make([]int32, len(intArr))
	copy(target, intArr)
	// 中間
	mid := int32((len(target) - 1) / 2)

	for i := mid; i >= 0; i-- {
		maxHeapify(target, i)
	}
	return target
}

func maxHeapify(arr []int32, i int32) {
	left := 2*i + 1
	right := 2*i + 2
	heapSize := int32(len(arr))
	largest := i

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest)
	}
}

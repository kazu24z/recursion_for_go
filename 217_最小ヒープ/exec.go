package main

func buildMinHeap(intArr []int32) []int32 {
	result := make([]int32, len(intArr))
	copy(result, intArr)
	mid := int32(len(result)-1) / 2

	for i := mid; i >= 0; i-- {
		minHeapify(result, i)
	}
	return result
}

func minHeapify(arr []int32, i int32) {
	left := 2*i + 1
	right := 2*i + 2
	heapSize := len(arr)
	largest := i

	if left < int32(heapSize) && arr[left] < arr[largest] {
		largest = left
	}

	if right < int32(heapSize) && arr[right] < arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		minHeapify(arr, largest)
	}

}

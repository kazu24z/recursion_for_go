package main

import (
	"fmt"
)

func main() {
	arr := []int32{2, 42, 11, 30, 10, 7, 6, 5, 9}
	heap := Heap{}

	fmt.Println("Before:", arr)
	heap.HeapSort(arr)
	fmt.Println("After:", arr)
}

func (h *Heap) HeapSort(arr []int32) {
	// 最大ヒープ化する
	h.BuildMaxHeap(arr)

	// 配列の範囲を定義
	arrRange := len(arr) - 1

	// 配列範囲が0より大きい限り処理を繰り返す
	for arrRange > 0 {
		// 配列先頭と末尾を入れ替える
		temp := arr[0]
		arr[0] = arr[arrRange] // 先頭に末尾の値を入れる
		arr[arrRange] = temp   // 末尾に先頭の値を入れる

		// 入れ替えたら最大ヒープが崩れるから最大ヒープ化
		// 配列範囲を1つ狭める
		h.MaxHeapify(arr, 0, int32(arrRange))
		arrRange--

	}
}

func (h *Heap) BuildMaxHeap(arr []int32) {
	mid := int32((len(arr) - 1) / 2)

	for i := mid; i >= 0; i-- {
		h.MaxHeapify(arr, i, int32(len(arr)))
	}

}

type Heap struct{}

func (h *Heap) MaxHeapify(arr []int32, i int32, heapRange int32) {
	left := h.left(i)
	right := h.right(i)

	largest := i

	if left < heapRange && arr[largest] < arr[left] {
		largest = left
	}

	if right < heapRange && arr[largest] < arr[right] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		h.MaxHeapify(arr, largest, heapRange)
	}

}

func (h *Heap) left(i int32) int32 {
	return 2*i + 1
}

func (h *Heap) right(i int32) int32 {
	return 2*i + 2
}

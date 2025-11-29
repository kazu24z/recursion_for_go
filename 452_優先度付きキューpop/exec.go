package main

import (
	"fmt"
)

func main() {

	pq := NewPriorityQueue([]int32{2, 3, 43, 2, 53, 6, 75, 10})
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)
	fmt.Println(pq.Pop())
	fmt.Println(pq.data)

}

type PriorityQueue struct {
	data []int32
}

func NewPriorityQueue(arr []int32) *PriorityQueue {
	// 元配列をコピー
	target := make([]int32, len(arr))
	copy(target, arr)

	// 最大ヒープを生成する処理
	pq := &PriorityQueue{}
	pq.buildMaxHeap(target)
	pq.data = target
	return pq

}

func (pq *PriorityQueue) Pop() int32 {
	if len(pq.data) <= 0 {
		panic("priority queue is empty")
	}
	// ルートが最大値なので、それを一時変数に入れる
	poped := pq.data[0]
	// 配列の末尾を先頭に入れる
	pq.data[0] = pq.data[len(pq.data)-1]

	// 配列の末尾から値を削除
	pq.data = pq.data[:len(pq.data)-1]

	// 先頭から再度ヒープ化処理を実行する
	pq.maxHeapify(pq.data, 0)
	return poped
}

func (pq *PriorityQueue) buildMaxHeap(arr []int32) {
	// 入力配列の真ん中
	mid := int32((len(arr) - 1) / 2)

	for i := mid; i >= 0; i-- {
		pq.maxHeapify(arr, i)
	}
}

func (pq *PriorityQueue) maxHeapify(arr []int32, i int32) {
	left := pq.left(i)
	right := pq.right(i)
	heapSize := int32(len(arr))
	largest := i

	// 左右と比較して左右が大きいならindexを入れ替える
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}

	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		// 値を入れ替える
		arr[i], arr[largest] = arr[largest], arr[i]
		pq.maxHeapify(arr, largest)
	}
}

func (pq *PriorityQueue) left(i int32) int32 {
	return 2*i + 1
}

func (pq *PriorityQueue) right(i int32) int32 {
	return 2*i + 2
}

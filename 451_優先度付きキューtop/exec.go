package main

import (
	"fmt"
)

func main() {

	pq1 := NewPriorityQueue([]int32{2, 3, 43, 2, 53, 6, 75, 10})
	fmt.Println(pq1.top())

	pq2 := NewPriorityQueue([]int32{3, 12, 0, 2, 9, 1, 65, 32})
	fmt.Println(pq2.top())

	pq3 := NewPriorityQueue([]int32{1, 2, 3, 4, 8, 2, 1, 9, 7, 3, 4})
	fmt.Println(pq3.top())

}

type PriorityQueue struct {
	maxHeap []int32
}

func NewPriorityQueue(arr []int32) *PriorityQueue {
	// 入力配列のコピーを作成
	target := make([]int32, len(arr))
	copy(target, arr)

	// 最大ヒープ化
	heapLibrary := HeapLibrary{}
	heapLibrary.buildMaxHeap(target)

	return &PriorityQueue{maxHeap: target}
}

func (pq *PriorityQueue) top() int32 {
	if len(pq.maxHeap) == 0 {
		// エラーハンドリング（パニックまたはデフォルト値）
		panic("empty priority queue")
		// または return 0 など
	}
	return pq.maxHeap[0]
}

type HeapLibrary struct {
	//
}

// 全体を最大ヒープにする
func (hl *HeapLibrary) buildMaxHeap(arr []int32) {
	// 配列全体の真ん中
	mid := int32((len(arr) - 1) / 2) // 6個なら[2]

	// 真ん中から前に向かってヒープ処理を繰り返す
	for i := mid; i >= 0; i-- {
		hl.maxHeapify(arr, i)
	}
}

// 親子関係ででかい方が親になるやつ（入力以下再帰的にチェック）
func (hl *HeapLibrary) maxHeapify(arr []int32, i int32) {
	left := hl.left(i)
	right := hl.right(i)
	largest := i

	heapSize := int32(len(arr))

	if left < heapSize && arr[largest] < arr[left] {
		largest = left
	}

	if right < heapSize && arr[largest] < arr[right] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		hl.maxHeapify(arr, largest)
	}

}

func (hl *HeapLibrary) left(i int32) int32 {
	return 2*i + 1
}

func (hl *HeapLibrary) right(i int32) int32 {
	return 2*i + 2
}

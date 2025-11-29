package main

import (
	"fmt"
)

func main() {
	arr := []int32{2, 42, 11, 30, 10, 7, 6, 5, 9}
	heap := Heap{}

	fmt.Println("=== ヒープソート開始 ===")
	fmt.Println("初期配列:", arr)
	heap.HeapSort(arr)
	fmt.Println("=== ヒープソート完了 ===")
	fmt.Println("最終結果:", arr)
}

func (h *Heap) HeapSort(arr []int32) {
	fmt.Println("\n--- 最大ヒープ構築開始 ---")
	// 最大ヒープ化する
	h.BuildMaxHeap(arr)
	fmt.Println("最大ヒープ構築完了:", arr)

	// 配列の範囲を定義
	arrRange := len(arr) - 1

	fmt.Println("\n--- ソート処理開始 ---")
	// 配列範囲が0より大きい限り処理を繰り返す
	for arrRange > 0 {
		fmt.Printf("\n[ステップ %d] 範囲: 0〜%d\n", len(arr)-arrRange, arrRange)
		fmt.Printf("  交換前: %v (最大値 %d を末尾へ)\n", arr, arr[0])

		// 配列先頭と末尾を入れ替える
		temp := arr[0]
		arr[0] = arr[arrRange] // 先頭に末尾の値を入れる
		arr[arrRange] = temp   // 末尾に先頭の値を入れる

		fmt.Printf("  交換後: %v\n", arr)
		fmt.Printf("  確定済み: %v\n", arr[arrRange:])

		// 入れ替えたら最大ヒープが崩れるから最大ヒープ化
		// 配列範囲を1つ狭める
		fmt.Printf("  ヒープ修復対象範囲: 0〜%d\n", arrRange-1)
		h.MaxHeapify(arr, 0, int32(arrRange))
		fmt.Printf("  ヒープ修復後: %v\n", arr[:arrRange])
		arrRange--

	}
}

func (h *Heap) BuildMaxHeap(arr []int32) {
	mid := int32((len(arr) - 1) / 2)
	fmt.Printf("最後の非葉ノード: インデックス %d\n", mid)

	for i := mid; i >= 0; i-- {
		fmt.Printf("  インデックス %d からMaxHeapify実行\n", i)
		h.MaxHeapify(arr, i, int32(len(arr)))
		fmt.Printf("    結果: %v\n", arr)
	}

}

type Heap struct{}

func (h *Heap) MaxHeapify(arr []int32, i int32, heapRange int32) {
	left := h.left(i)
	right := h.right(i)

	largest := i

	fmt.Printf("    MaxHeapify(i=%d): 親=%d", i, arr[i])

	if left < heapRange {
		fmt.Printf(", 左子=%d", arr[left])
		if arr[largest] < arr[left] {
			largest = left
		}
	} else {
		fmt.Printf(", 左子=なし")
	}

	if right < heapRange {
		fmt.Printf(", 右子=%d", arr[right])
		if arr[largest] < arr[right] {
			largest = right
		}
	} else {
		fmt.Printf(", 右子=なし")
	}

	if largest != i {
		fmt.Printf(" → 交換: %d ↔ %d\n", arr[i], arr[largest])
		arr[i], arr[largest] = arr[largest], arr[i]
		h.MaxHeapify(arr, largest, heapRange)
	} else {
		fmt.Printf(" → 交換なし\n")
	}

}

func (h *Heap) left(i int32) int32 {
	return 2*i + 1
}

func (h *Heap) right(i int32) int32 {
	return 2*i + 2
}

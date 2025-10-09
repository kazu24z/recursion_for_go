package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func findMinNum(head *SinglyLinkedListNode) int32 {
	min := head.data
	counter := int32(0)
	targetIndex := int32(0)

	for current := head; current != nil; current = current.next {
		if min >= current.data {
			min = current.data
			targetIndex = counter
		}
		counter++
	}

	return targetIndex
}

// 片方向リストなので線形探索するが、その過程で最小を更新し続ける流れ

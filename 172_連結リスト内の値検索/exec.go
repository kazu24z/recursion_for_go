package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func linkedListSearch(head *SinglyLinkedListNode, data int32) int32 {
	index := int32(0)
	resultIndex := int32(-1)
	for current := head; current != nil; current = current.next {
		if current.data == data {
			resultIndex = index
			break // 一致した時点でこのループ不要なので抜ける
		}

		index++
	}

	return resultIndex
}

// ループでheadからdataを見て最初にヒットした回のindex(0始まり)を返す

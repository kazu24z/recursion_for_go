package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertHeadTail(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	// 先頭に挿入
	newHead := &SinglyLinkedListNode{
		data: data,
		next: head,
	}

	// 元のheadから末尾まで移動
	current := head
	for current.next != nil {
		current = current.next
	}

	// 末尾に挿入
	current.next = &SinglyLinkedListNode{
		data: data,
		next: nil,
	}

	return newHead
}

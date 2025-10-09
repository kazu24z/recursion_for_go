package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func linkedListLength(head *SinglyLinkedListNode) int32 {
	counter := int32(0)
	current := head
	for current != nil {
		current = current.next
		counter++
	}

	return counter
}

// カウンターを使ってループ回数を数える

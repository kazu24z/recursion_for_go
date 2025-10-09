package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func linkedListLastValue(head *SinglyLinkedListNode) int32 {
	// ループを辿ってnextがnilのnodeのdataを返す
	current := head
	for current.next != nil {
		current = current.next
	}

	return current.data
}

package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func getIndexValue(head *SinglyLinkedListNode, index int32) int32 {
	currentNode := head

	for i := int32(1); i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.data
}

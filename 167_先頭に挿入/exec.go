package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertAtHead(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	// あたらしいNodeを用意
	targetNode := &SinglyLinkedListNode{
		data: data,
		next: head,
	}

	// 上記をreturn
	return targetNode
}

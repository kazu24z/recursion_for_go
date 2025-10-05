package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertAtTail(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	// 概要：headから末尾を探す処理

	// 新しいnodeを作成
	newNode := &SinglyLinkedListNode{
		data: data,
		next: nil,
	}

	// 末尾のnodeのnextに新しいnodeを入れる
	currentNode := head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	// ループで末尾nodeが見つかるのでそのnextにnewNodeを代入
	currentNode.next = newNode

	return head
}

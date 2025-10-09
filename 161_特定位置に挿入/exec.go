package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeInSorted(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{
		data: data,
	}

	current := head
	// 先頭に挿入する場合
	if data <= current.data {
		newNode.next = current
		return newNode
	}

	for current.next != nil {
		// 挿入が間にある場合
		if current.data <= data && data <= current.next.data {
			// 挿入処理
			temp := current.next
			current.next = &SinglyLinkedListNode{
				data: data,
				next: temp,
			}
			return head
		}
		current = current.next
	}

	current.next = newNode

	return head
}

// 前後の値を保持しておく必要がある？？
// → currentとcurrent.next？

package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertAtPosition(head *SinglyLinkedListNode, position int32, data int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{data: data}

	current := head
	counter := int32(0)

	// position番目のノードまで移動
	for counter < position && current != nil {
		current = current.next
		counter++
	}

	// position番目のノードの後に挿入
	if current != nil {
		newNode.next = current.next
		current.next = newNode
	}

	return head
}

// position番目のnodeのnextをdataを持った新しいnodeにし、そのnextにもともとのnextを入れる

// まずポインタを移動→止まって挿入処理 という流れで良かった。横着した

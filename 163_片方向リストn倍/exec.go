package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func reproduceByN(head *SinglyLinkedListNode, n int32) *SinglyLinkedListNode {

	// 倍化の回数をカウント
	counter := int32(1)
	current := head
	copyTarget := deepCopyList(head)

	for current != nil {
		// 末尾がnil
		if current.next == nil && counter < n {
			// 元の片方向リストのコピーを渡さないとだめ
			current.next = deepCopyList(copyTarget)
			counter++
		}
		current = current.next
	}

	return head
}

func deepCopyList(originHead *SinglyLinkedListNode) *SinglyLinkedListNode {
	if originHead == nil {
		return nil
	}

	var newHead *SinglyLinkedListNode
	var newTail *SinglyLinkedListNode
	current := originHead

	for current != nil {
		newNode := &SinglyLinkedListNode{data: current.data}

		if newHead == nil {
			// リストの初期化
			newHead = newNode
			newTail = newNode // newHead と newTail は同じノードを指す
		} else {
			// 既存の末尾ノードの Next に新しいノードを連結
			newTail.next = newNode

			// 連結後、末尾ポインタを新しいノードに更新
			newTail = newNode
		}

		current = current.next
	}
	return newHead
}

// n倍の個数に増やす
// 各リストのtailノードのnextに次のheadを渡す

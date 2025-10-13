package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func palindromeLinkedList(head *SinglyLinkedListNode) bool {
	// ノード数をカウント
	current := head
	counter := int32(0)
	for current != nil {
		current = current.next
		counter++
	}

	temp := head
	stack := []int32{}

	// 前半をスタックに積む
	for i := int32(0); i < counter/2; i++ {
		stack = append(stack, temp.data)
		temp = temp.next
	}

	// 奇数の場合は真ん中をスキップ
	if counter%2 == 1 {
		temp = temp.next
	}

	// 後半をスタックと比較
	for i := int32(0); i < counter/2; i++ {
		if len(stack) == 0 || stack[len(stack)-1] != temp.data {
			return false
		}
		stack = stack[:len(stack)-1]
		temp = temp.next
	}

	return len(stack) == 0
}

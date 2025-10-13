package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// スロー/ファストポインタパターン
func palindromeLinkedList(head *SinglyLinkedListNode) bool {
	// 要素が1つの場合は回文扱い
	if head.next == nil {
		return true
	}

	// headを指す2つのポインタ
	slow := head
	fast := head

	stack := []int32{}

	// fastが最後に行くまで進める
	for fast != nil && fast.next != nil {
		// スタックにslow.dataを格納
		stack = append(stack, slow.data)

		slow = slow.next
		fast = fast.next.next
	}

	// fastがnil、つまり奇数の場合
	// 奇数の場合、末尾までfastが行くと、最後のfast = fast.nextでnilが入る
	if fast != nil {
		slow = slow.next
	}

	// slowには残り半分のnodeがリストになっている
	for slow != nil {
		// slow.dataと、stackの先頭（=配列の末尾）が一致していなかったらその時点でfalse
		if slow.data != stack[len(stack)-1] {
			return false
		}
		// stackの末尾をpop
		stack = stack[:len(stack)-1]
		slow = slow.next
	}

	return true
}

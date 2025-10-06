package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func mergeTwoSortedLinkedLists(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	// ダミーヘッド
	dummyHead := &SinglyLinkedListNode{}
	current := dummyHead // currentポインタを追加

	current1 := head1
	current2 := head2

	for current1 != nil && current2 != nil {
		if current1.data <= current2.data {
			current.next = current1
			current1 = current1.next
		} else {
			current.next = current2
			current2 = current2.next
		}
		current = current.next // currentを次に進める
	}

	// 残りを接続
	if current1 != nil {
		current.next = current1
	} else {
		current.next = current2
	}

	return dummyHead.next
}

// ソート済みなので、head1 vs head2で一騎打ちして小さい方を採用していく方向性
// 片方がなくなったらそのままnextに残りをくっつける

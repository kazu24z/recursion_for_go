package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func removeNthNode(head *SinglyLinkedListNode, n int32) *SinglyLinkedListNode {
	// headからtailまでのnode数を辿る
	nodeNum := int32(0)
	current := head
	for current != nil {
		current = current.next
		nodeNum++
	}

	// 前から何番目のノードを削除するか
	nodeNumFromFront := nodeNum - n + 1

	// nが0、またはリストの範囲外の場合はそのまま返す
	if n == 0 || nodeNumFromFront <= 0 || nodeNumFromFront > nodeNum {
		return head
	}

	// 先頭ノードを削除する場合
	if nodeNumFromFront == 1 {
		return head.next
	}

	// 削除するノードの1つ前のノードを見つける
	current = head
	for i := int32(1); i < nodeNumFromFront-1; i++ {
		current = current.next
	}

	// 削除するノードをスキップして繋ぎ変える
	current.next = current.next.next

	return head
}

// 要するに該当ノードを削除して、前後をつなぎ合わせる操作
// 末尾から辿る
// 1. 末尾までを１度走査して全体のノード数を把握
// 2. ノード数 - 指定数 + 1 = 前から辿った数
// 3. nが0以下またはリストの範囲外 → return head
// 4. 先頭ノードを削除する場合 → return head.next
// 5. 削除するノードの1つ前を見つけて、そのnextを削除ノードのnextに繋ぐ

// ポインタのイメージがまだちゃんとできてない。

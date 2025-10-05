package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func middleNode(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	counter := 1

	current := head
	for current.next != nil {
		current = current.next
		counter++
	}

	targetNum := counter/2 + 1 // 奇数も偶数も商+1がターゲット

	// 最後のnodeをheadとする新しい片方向リストを作成
	targetNode := head
	for i := 1; i < targetNum; i++ {
		targetNode = targetNode.next
	}

	return targetNode
}

// headしかわからない状況→全部走査して何個のnodeがあるのかを把握

// 奇数→node数/2の商+1番目のnodeが対象
// 偶数→node数/2の商+1番目のnodeが対象（間じゃなくて2つめの方だから）
// → 商+1番目のnode

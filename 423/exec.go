package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func doubleEvenNumber(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// 関数を完成させてください
	// 現在操作node
	currentNode := head
	// カウンタ
	index := 0
	for currentNode != nil {
		if index%2 == 0 {
			newNode := &SinglyLinkedListNode{
				data: currentNode.data * 2,
				next: currentNode.next,
			}

			// 連結
			currentNode.next = newNode

			// 次の回のnode
			currentNode = newNode.next // こっちで追加した分はスキップできている

			index++ // だから、こっちも+1するだけでいい
		} else {
			currentNode = currentNode.next
			// カウンタを1加算
			index++
		}
	}

	return head
}

// indexで偶数目のとき（0, 2, 4...）をピックアップする

// その回のノードに対しての処理
//// 1. 旧nextのnodeをtempへ
//// 2. 新node（dataを2倍したやつ）作成
//// 3. ノードのnextに新node, 新nodeのnextに旧node tempを代入

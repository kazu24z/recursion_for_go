package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func deleteTail(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	nodeNum := 0
	current := head

	for current.next != nil {
		current = current.next
		nodeNum++
	}

	current = head

	// node数が1つの場合
	if nodeNum <= 1 {
		return nil
	}

	for i := 0; i < nodeNum-1; i++ {
		current = current.next
	}

	// 最後から1つ手前のnodeがcurrentに入っているはずなので
	current.next = nil

	return head
}

// 末尾を特定
// 末尾の1つ前のnextをnilにするだけ
// 全体のnode数を特定
// 1つ前まででループを止める

/**
  // リストが空またはノードが1つしかない場合
  if head == nil || head.next == nil {
      return nil
  }

  current := head
  // 末尾の1つ前のノードを見つける
  for current.next.next != nil {
      current = current.next
  }

  // 最後のノードを削除
  current.next = nil

  return head


*/

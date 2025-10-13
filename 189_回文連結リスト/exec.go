package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func palindromeLinkedList(head *SinglyLinkedListNode) bool {
	// ノードの数
	current := head
	counter := int32(0)
	for current != nil {
		current = current.next
		counter++
	}

	isEven := (counter%2 == 0) // 偶数ならtrue
	var skipNum int32

	if isEven {
		skipNum = 1
	} else {
		skipNum = 2
	}

	temp := head
	// 空のスタック配列
	stack := []int32{}

	// 先頭からcounterの半分までをスタックに入れる
	for i := int32(1); i <= counter/2; i++ {
		stack = append(stack, temp.data)
		temp = temp.next
	}

	if !isEven {
		// 奇数の場合、頂点の値をスキップ
		temp = temp.next
	}

	// 後半分に対して、スタックの先頭と比較し、値が一致し続けるかをチェック
	for i := counter/2 + skipNum; i <= counter; i++ {
		if stack[len(stack)-1] == temp.data {
			// スタックから値をpopしていく
			stack = stack[:len(stack)-1]
		}
		temp = temp.next
	}
	// スタックが空なら回文
	if len(stack) == 0 {
		return true
	}

	// それ以外の場合はfalseを返す
	return false
}

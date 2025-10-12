package main

func consecutiveWalk(arr []int32) []int32 {
	// 配列が与えられて、その中から単調増加している部分のうち、最後にそうなっている部分を出力する
	// 条件を満たしたものをStackに積んでいく
	//// →単調増加部分
	//// →Stackの先頭より小さい値がきたら、それまでのStackをPopして、結果配列に入れる

	stack := &Stack{} // 抽出したものを入れていくStack

	for i := 0; i < len(arr); i++ {
		current := arr[i]
		// stackがnilまたは、stackの頭の値がその回の配列要素値より小さいなら、Stackの先頭にその回の値を追加
		if stack.head == nil || stack.peek() < current {
			stack.push(current)
		} else {
			// 単調増加が途切れたら、stackをクリア
			for stack.head != nil {
				stack.pop()
			}
			stack.push(current)
		}

		current = arr[i]
	}

	resultArr := make([]int32, 0) // 最終的な結果を入れるスライス
	for stack.head != nil {
		value := stack.pop()
		if value != nil {
			resultArr = append(resultArr, *value)
		}
	}

	return resultArr
}

type Item struct {
	data int32
	next *Item
}

type Stack struct {
	head *Item
}

// 先頭に新しいnodeを追加
func (s *Stack) push(data int32) {
	// 現在のheadを一時退避
	temp := s.head
	// 新しいnodeを作成
	// 新しいnodeをheadに格納
	s.head = &Item{data: data}

	// 新しいheadノードのnextに避難したnodeを格納
	s.head.next = temp
}

// 先頭のnodeの値を取り出す
func (s *Stack) pop() *int32 {
	if s.head == nil {
		return nil
	}

	data := s.head.data
	s.head = s.head.next

	return &data
}

// 先頭のnodeの値を返す
func (s *Stack) peek() int32 {
	data := s.head.data

	return data
}

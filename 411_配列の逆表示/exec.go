package main

import "errors"

func reverse(arr []int32) []int32 {
	length := len(arr)

	// 空のstack
	stack := &Stack{}

	// arr要素分push
	for i := 0; i < length; i++ {
		stack.push(arr[i])
	}

	// スライス
	reversed := make([]int32, length)
	// stackのheadから末尾までpop
	for i := 0; i < length; i++ {
		data, err := stack.pop()
		if err != nil {
			// エラー処理
		} else {
			// 新しいスライスに追加
			reversed[i] = data
		}
	}

	return reversed
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

var ErrStackEmpty = errors.New("stack is empty")

// 先頭のnodeを取り出す
func (s *Stack) pop() (int32, error) {
	if s.head == nil {
		// スタックが空の場合は、intのゼロ値とfalse（失敗）を返す
		return 0, ErrStackEmpty
	}

	data := s.head.data
	s.head = s.head.next

	// データ（0を含む）とtrue（成功）を返す
	return data, nil
}

// 先頭のnodeの値を返す
func (s *Stack) peek() int32 {
	data := s.head.data

	return data
}

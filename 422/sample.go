package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 5, 3, 7}

	list := NewSingleLinkedList(arr)

	test1 := list.At((2))
	fmt.Printf("Indexが2のNodeの値は%dです。\n", test1.Data)

	test2 := list.FindNode(5)
	fmt.Printf("値が%dなNodeのIndexは%dです。\n", 5, test2)
}

// ノード
type Node struct {
	Data int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

// 片方向リスト
type SingleLinkedList struct {
	Head *Node
}

func NewSingleLinkedList(arr []int) *SingleLinkedList {
	// Headを確定
	headNode := NewNode(arr[0]) // Nodeのポインタインスタンス = 以降実体へのアクセス

	// 上のheadに対して、連鎖的にnextに入れていく
	currentNode := headNode // これは現在操作中のnodeのアドレスが入る

	for i := 1; i < len(arr); i++ {
		currentNode.Next = NewNode(arr[i]) // 現在操作中のアドレスから実体をデリファレンスし、そこに新しいNdeインスタンスのアドレスが入ってる
		currentNode = currentNode.Next     // 次のNodeを操作対象にセット
	}

	return &SingleLinkedList{Head: headNode}
}

func (s *SingleLinkedList) At(index int) *Node {
	// 受け取ったindex回分ループして、その回目にあるNodeを取り出す
	// 先頭Nodeを取り出す
	target := s.Head

	for i := 0; target.Next != nil && i < index; i++ {
		target = target.Next
	}

	return target
}

func (s *SingleLinkedList) FindNode(key int) int {
	target := s.Head
	counter := 0
	for i := 0; target != nil; i++ {
		if target.Data == key {
			counter = i
		}
		target = target.Next
	}

	return counter
}

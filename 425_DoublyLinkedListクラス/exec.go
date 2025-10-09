package main

import (
	"fmt"
)

func main() {
	numList := NewDoublyLinkedList([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(numList.head.data)
	fmt.Println(numList.head.next.data)
	fmt.Println(numList.head.next.prev.data)
	fmt.Println(numList.tail.data)
	fmt.Println(numList.tail.prev.data)
	fmt.Println(numList.tail.prev.prev.data)
}

type Item struct {
	data int
	prev *Item
	next *Item
}

type DoublyLinkedList struct {
	head *Item
	tail *Item
}

func NewDoublyLinkedList(arr []int) *DoublyLinkedList {
	// この中でItemの生成→それぞれを双方向リストになるように連結する処理を書いていく

	// イメージ：headの生成以降でnextやprevに格納していく

	// 空配列の場合、headもtailのnilのDoublyLinkedListインスタンスを返す
	if len(arr) == 0 {
		return &DoublyLinkedList{}
	}

	head := &Item{
		data: arr[0],
	}

	current := head

	// arrの要素数分繰り返す
	for i := 1; i < len(arr); i++ {
		current.next = &Item{
			data: arr[i],
			prev: current,
		}

		current = current.next
	}

	return &DoublyLinkedList{
		head: head,
		tail: current,
	}
}

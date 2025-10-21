package main

import "fmt"

func main() {
	q := NewQueue()
	q.Enqueue(4)

	fmt.Println(q.PeekFront())

	fmt.Println(q.PeekBack())

	q.Enqueue(50)

	fmt.Println(q.PeekFront())

	fmt.Println(q.PeekBack())

	q.Enqueue(64)

	fmt.Println(q.PeekFront())

	fmt.Println(q.PeekBack())

	fmt.Println(q.Dequeue())
}

type Node struct {
	data int
	next *Node
}

/**
--------------------------------
tail 	 					head →
--------------------------------
*/

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil}
}

func (q *Queue) PeekFront() int {
	return q.head.data
}

func (q *Queue) PeekBack() int {
	return q.tail.data
}

func (q *Queue) Enqueue(data int) {
	// 追加するノード
	newNode := &Node{data: data}

	temp := q.tail     // 現在の末尾を退避
	q.tail = newNode   // 現在の末尾に新しいノードを格納
	q.tail.next = temp // 新しいノードのnextに古い末尾を格納

	// キューが空の場合に追加すると、tail = headになるため
	if q.head == nil {
		q.head = newNode
	}
}

func (q *Queue) Dequeue() int {
	// headの1つ前を探る => tailから辿る
	current := q.tail
	for current.next != nil {
		current = current.next
	}

	result := current.data

	// それをheadにする
	q.head = current
	// 新しいheadのnextをnilにする
	q.head.next = nil

	return result
}

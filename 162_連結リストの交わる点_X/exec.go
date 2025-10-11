package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func findMergeNode(headA *SinglyLinkedListNode, headB *SinglyLinkedListNode) int32 {
	reversedA := reverseNode(headA)
	reversedB := reverseNode(headB)

	// 先頭から連続で一致するnodeがあるかをチェック
	currentA := reversedA
	currentB := reversedB

	result := int32(-1)

	for currentA.data == currentB.data {
		fmt.Println("currentA.data:", currentA.data)
		fmt.Println("currentB.data:", currentB.data)
		if currentA != nil && currentB != nil {
			break
		}
		result = currentA.data
		fmt.Println("result", result)
		currentA = currentA.next
		currentB = currentB.next

	}
	fmt.Println("result", result)
	printNodeList(reversedA, "reversedA")
	printNodeList(reversedB, "reversedB")

	return result
}

func reverseNode(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// １つ前、headから見たらnilになるから初期値nil
	var prev *SinglyLinkedListNode = nil
	current := head

	for current != nil {
		// 自分から見た次を退避させる
		next := current.next
		// 自分の次を１つ前に反転
		current.next = prev
		// 次のループのために、１つ前に自分自身をセット
		prev = current
		current = next
	}

	return prev
}

func printNodeList(head *SinglyLinkedListNode, name string) {
	fmt.Println(name)
	for current := head; current != nil; current = current.next {
		if current.next == nil {
			fmt.Println(current.data, "→END")
		} else {
			fmt.Print(current.data, "→")
		}
	}
}

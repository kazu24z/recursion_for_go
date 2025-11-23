package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// テストデータ
	arr := []int32{10, 5, 15, 3, 7, 12, 20}

	bst := &BinarySearchTree{}
	result := bst.generateRandomBST(arr)

	fmt.Println("生成されたBSTの間順走査（ソート順）:")
	result.printInOrder(result.root)
	fmt.Println()

	fmt.Println("生成されたBSTの前順走査（構造確認）:")
	result.printPreOrder(result.root)
}

// 間順走査（ソート順で表示）
func (b *BinarySearchTree) printInOrder(node *BinaryTree) {
	if node == nil {
		return
	}
	b.printInOrder(node.left)
	fmt.Printf("%d ", node.data)
	b.printInOrder(node.right)
}

// 前順走査（木の構造確認用）
func (b *BinarySearchTree) printPreOrder(node *BinaryTree) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.data)
	b.printPreOrder(node.left)
	b.printPreOrder(node.right)
}

type BinaryTree struct {
	data  int32
	left  *BinaryTree
	right *BinaryTree
}

type BinarySearchTree struct {
	root *BinaryTree
}

func (b *BinarySearchTree) Insert(v int32) {
	newNode := &BinaryTree{data: v}
	if b.root == nil {
		b.root = &BinaryTree{data: v}
		return
	}

	iterator := b.root

	for iterator != nil {
		if iterator.data > v {
			// もしiterator.leftがnilになるなら、そのleftにnewNode入れてbreak
			if iterator.left == nil {
				iterator.left = newNode
				break
			}

			iterator = iterator.left
		} else {
			if iterator.right == nil {
				iterator.right = newNode
				break
			}
			iterator = iterator.right
		}
	}
}

func (b *BinarySearchTree) shuffle(arr []int32) []int32 {
	num := len(arr)

	indexArr := make([]int32, 0)
	newArr := make([]int32, 0)

	// 0からnum-1までの正の整数 i を生成する処理
	for i := 0; i < num; i++ {
		val := int32(rand.Intn(num)) // indexを生成
		if !contains(indexArr, val) {
			indexArr = append(indexArr, val)
			newArr = append(newArr, arr[val])
		} else {
			i--
		}
	}
	return newArr
}

func (b *BinarySearchTree) generateRandomBST(arr []int32) *BinarySearchTree {
	if len(arr) == 0 {
		return &BinarySearchTree{root: nil}
	}

	shuffled := b.shuffle(arr)

	bst := &BinarySearchTree{}
	for i := 0; i < len(shuffled); i++ {
		bst.Insert(shuffled[i])
	}
	return bst
}

// ===== Util ====
func contains(arr []int32, v int32) bool {
	for _, x := range arr {
		if x == v {
			return true
		}
	}
	return false
}

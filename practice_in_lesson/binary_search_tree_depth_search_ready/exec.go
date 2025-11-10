package main

import (
	"fmt"
)

func main() {
	arr := []int32{4, 43, 36, 46, 32, 7, 97, 95, 34, 8, 96, 35, 85, 1010, 232}

	bst := NewBinarySearchTree(arr)

	fmt.Println("key97は存在してる？:", bst.KeyExist(97))
	fmt.Println("keyが97:", bst.Search(97).data)

	fmt.Print("間順走査 結果： ")
	bst.root.PrintInOrder()
	fmt.Print("逆換順走査 結果: ")
	bst.root.PrintReverseOrder()
	fmt.Print("前順走査 結果: ")
	bst.root.PrintPreOrder()
	fmt.Print("後順走査 結果: ")
	bst.root.PrintPostOrder()
}

// BinaryTree = 単なる木構造
type BinaryTree struct {
	data  int32
	left  *BinaryTree
	right *BinaryTree
}

// 間順走査の結果を出力する関数
func (b *BinaryTree) PrintInOrder() {
	// memo: 間順 → 左→自分→右 の順での走査
	b.inOrderWalk(b)
	fmt.Println()
}

func (b *BinaryTree) PrintReverseOrder() {
	b.reverseOrder(b)
	fmt.Println()
}

func (b *BinaryTree) PrintPreOrder() {
	b.preOrderWalk(b)
	fmt.Println()
}

func (b *BinaryTree) PrintPostOrder() {
	b.postOrderWalk(b)
	fmt.Println()
}

// 間順走査
func (b *BinaryTree) inOrderWalk(root *BinaryTree) {
	if root != nil {
		b.inOrderWalk(root.left)
		fmt.Print(root.data, " ")
		b.inOrderWalk(root.right)
	}
}

// 逆間順走査
// 右→自分→左の再帰処理で降順に取り出す
func (b *BinaryTree) reverseOrder(root *BinaryTree) {
	if root != nil {
		b.reverseOrder(root.right)
		fmt.Print(root.data, " ")
		b.reverseOrder(root.left)
	}
}

// 前順走査
// 自分→左→右の再帰処理で取り出す（木構造の状態を保持して取り出せる）
func (b *BinaryTree) preOrderWalk(root *BinaryTree) {
	if root != nil {
		fmt.Print(root.data, " ")
		b.preOrderWalk(root.left)
		b.preOrderWalk(root.right)
	}
}

// 後順走査
// 左→右→自分の再帰処理
func (b *BinaryTree) postOrderWalk(root *BinaryTree) {
	if root != nil {
		b.postOrderWalk(root.left)
		b.postOrderWalk(root.right)
		fmt.Print(root.data, " ")

	}
}

type BinarySearchTree struct {
	root *BinaryTree
}

func NewBinarySearchTree(arr []int32) *BinarySearchTree {
	// 配列をソートする
	sortedArr := sortArray(arr)

	return &BinarySearchTree{root: sortedArrayToBST(sortedArr)}
}

// keyから対象のnodeを見つける
func (b *BinarySearchTree) Search(key int32) *BinaryTree {
	iterator := b.root
	for iterator != nil {
		if key == iterator.data {
			return iterator
		}

		if key < iterator.data {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return nil
}

func (b *BinarySearchTree) KeyExist(key int32) bool {
	iterator := b.root
	for iterator != nil {
		if key == iterator.data {
			return true
		}

		if key < iterator.data {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return false
}

// ============ Utilities =====================================
func sortArray(arr []int32) []int32 {
	// マージソート
	return sortArrayHelper(arr, 0, len(arr)-1)
}

func sortArrayHelper(arr []int32, start int, end int) []int32 {
	if start == end {
		return []int32{arr[start]}
	}

	// ベースケース
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	left := sortArrayHelper(arr, start, mid-1)
	right := sortArrayHelper(arr, mid+1, end)

	// 勝ち抜け
	var combinedArr []int32
	i, j := 0, 0
	for len(left) > i && len(right) > j {
		if left[i] <= right[j] {
			combinedArr = append(combinedArr, left[i])
			i++
		} else {
			combinedArr = append(combinedArr, right[j])
			j++
		}
	}

	combinedArr = append(combinedArr, left[i:]...)
	combinedArr = append(combinedArr, right[j:]...)

	return combinedArr
}

func sortedArrayToBST(arr []int32) *BinaryTree {
	return sortedArrayToBSTHelper(arr, 0, len(arr)-1)
}

func sortedArrayToBSTHelper(arr []int32, start int, end int) *BinaryTree {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	left := sortedArrayToBSTHelper(arr, start, mid-1)
	right := sortedArrayToBSTHelper(arr, mid+1, end)

	root := &BinaryTree{data: arr[mid], left: left, right: right}
	return root
}

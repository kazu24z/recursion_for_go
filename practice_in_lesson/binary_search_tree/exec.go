package main

import "fmt"

func main() {
	arr := []int32{3, 6, 1, 4, 9, 2, 8, 7, 5}

	bst := NewBinarySearchTree(arr)
	fmt.Println(bst.root.data)
	fmt.Println(bst.root.left.data)
	fmt.Println(bst.root.right.data)
}

type BinaryTree struct {
	data  int32
	left  *BinaryTree
	right *BinaryTree
}

type BinarySearchTree struct {
	root *BinaryTree
}

func NewBinarySearchTree(arr []int32) *BinarySearchTree {
	// arrがソート済みとは限らないので、昇順ソートする
	sortedArr := mergeSortArr(arr)

	fmt.Println("ソート済み配列:", sortedArr)

	root := sortedArrayToBST(sortedArr)
	return &BinarySearchTree{root: root}
}

// ソート処理 せっかくなのでマージソート実装
func mergeSortArr(arr []int32) []int32 {
	return mergeSortArrHelper(arr, int32(0), int32(len(arr)-1))
}

func mergeSortArrHelper(arr []int32, start int32, end int32) []int32 {
	// ベースケース
	if start == end {
		return arr
	}

	// 分割統治法
	// 配列を2分割していく
	mid := int32((start + end) / 2)
	left := mergeSortArrHelper(arr[start:mid+1], 0, mid-start) // ベースケースを終えた後は[2]とかになっている = 要素数1のスライス
	right := mergeSortArrHelper(arr[mid+1:end+1], 0, end-mid-1)
	// memo: left, rightは再帰で、ここでは、とにかく分割していくことに徹する。
	// 最後まで分割した後で、結合処理をする。分けて考えること

	combined := make([]int32, 0)
	// left, rightが出揃った時点で両スライスで比較し、勝ち抜けで結果スライスに入れていく
	leftCounter := 0
	rightCounter := 0
	for len(left) > leftCounter && len(right) > rightCounter {
		if left[leftCounter] <= right[rightCounter] {
			combined = append(combined, left[leftCounter])
			leftCounter++
		} else {
			combined = append(combined, right[rightCounter])
			rightCounter++
		}
	}

	// 最後余ったほうを結果スライスのケツにappendしてreturn
	combined = append(combined, left[leftCounter:]...)
	combined = append(combined, right[rightCounter:]...)

	return combined
}

func sortedArrayToBST(arr []int32) *BinaryTree {
	return sortedArrayToBSTHelper(arr, 0, int32(len(arr)-1))
}

func sortedArrayToBSTHelper(arr []int32, start int32, end int32) *BinaryTree {
	// ベースケース
	if start > end {
		return nil
	}

	// 受け取ったスライスの中央
	mid := (start + end) / 2

	// midより前、後ろで再帰的に呼び出す
	left := sortedArrayToBSTHelper(arr, start, mid-1) // leftは常にstartは0
	right := sortedArrayToBSTHelper(arr, mid+1, end)

	// L, R両方とも帰ってきたら、BinaryTreeを組み立てる
	root := &BinaryTree{data: arr[mid], left: left, right: right} // このルートは各回=部分木のrootという意味

	return root
}

func (b *BinarySearchTree) Search(key int32) *BinaryTree {
	// BST構造体にこの関数を紐づかせるため、再帰じゃなくてrootからのイテレーションで

	// イテレーション
	for n := b.root; n != nil; {
		if n.data == key {
			return n
		}

		if n.data > key {
			n = n.left
		} else {
			n = n.right
		}
	}
	// 探索がうまくいく場合は上のイテレーションでreturnするためここに来たらハズレなのでnil返す
	return nil
}

func (b *BinarySearchTree) Exist(key int32) bool {
	// Searchと同様に探索して見つかったらtrue, 最後までなかったらfalse
	for n := b.root; n != nil; {
		if n.data == key {
			return true
		}

		if n.data > key {
			n = n.left
		} else {
			n = n.right
		}
	}

	return false
}

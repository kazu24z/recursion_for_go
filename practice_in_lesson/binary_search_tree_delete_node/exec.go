package main

import (
	"fmt"
)

// 二分木 > 削除処理
// 大きく4パターンあると

// 二分木からノードを特定する処理（search）

// 二分木削除処理

func main() {
	arr2 := []int32{2, 3, 6, 7, 8, 9, 10, 11, 12, 15}
	bst2 := NewBinarySearchTree(arr2)
	fmt.Println("ルート：", bst2.root.data)

	before2 := bst2.BstToArr()
	fmt.Println("削除前：", before2)

	targetNode2 := bst2.Search(3)
	bst2.DeleteNode(targetNode2)

	after2 := bst2.BstToArr()
	fmt.Println("削除後：", after2)
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
	sorted := mergeSort(arr) // ソート済みスライス
	root := makeArrToBSTHelper(sorted, 0, int32(len(sorted)-1))

	return &BinarySearchTree{root: root}
}

func makeArrToBSTHelper(arr []int32, start int32, end int32) *BinaryTree {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	left := makeArrToBSTHelper(arr, start, mid-1)
	right := makeArrToBSTHelper(arr, mid+1, end)

	return &BinaryTree{data: arr[mid], left: left, right: right}
}

// ノードを探す処理
func (b *BinarySearchTree) Search(key int32) *BinaryTree {
	iterator := b.root
	for iterator != nil {
		if iterator.data == key {
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

func (b *BinarySearchTree) IsExist(key int32) bool {
	iterator := b.root
	for iterator != nil {
		if iterator.data == key {
			return true
		}

		if iterator.data > key {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return false
}

// ノードを削除する処理
func (b *BinarySearchTree) DeleteNode(node *BinaryTree) {
	if !b.IsExist(node.data) {
		return
	}

	// ノードがない場合はそのまま処理を終わらせる
	if node == nil {
		return
	}

	// 対象がルートノードの場合
	if node == b.root {
		b.root = nil
		return
	}

	target, parent := b.getAncestorAndDescendant(node.data)
	fmt.Println("ターゲット：", target.data)
	if parent != nil {
		fmt.Println("親ノード　：", parent.data)
	}
	// 対象が葉の場合、それを削除するだけ
	if parent != nil && target.left == nil && target.right == nil {
		if target == parent.left {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}

	// 削除対象ノードの子が左側のみの場合、直下の子ノードを削除対象ノードの親ノードの子にする
	fmt.Println("左", target.left)
	fmt.Println("右", target.right)
	if parent != nil && target.left != nil && target.right == nil {
		tmp := target.left
		if target == parent.left {
			parent.left = tmp
		} else {
			parent.right = tmp
		}
		return
	}

	// 削除対象ノードの子が右側のみの場合、削除対象ノードの右側の子（部分木）を自分の場所に持っていく
	if parent != nil && target.right != nil && target.left == nil {
		if target == parent.left {
			parent.left = target.right
		} else {
			parent.right = target.right
		}
		return
	}

	// 削除対象ノードの左右に子がある場合、削除対象の右側部分木の最小値ノードが、自分の場所にくる
	if parent != nil && target.right != nil && target.left != nil {
		minNode, minNodeParent := b.getMinNodeWithParent(target.right)
		fmt.Println("部分木最小：", minNode)

		if minNodeParent != nil && minNodeParent != target {
			minNodeParent.left = minNode.right
			minNode.right = target.right
		}

		minNode.left = target.left
		if target == parent.left {
			parent.left = minNode
		} else {
			parent.right = minNode
		}
		return
	}

}

// 対象ノードとその親を返す
func (b *BinarySearchTree) getAncestorAndDescendant(key int32) (*BinaryTree, *BinaryTree) {
	//
	current := b.root
	var parent *BinaryTree
	for current != nil {
		if current.data == key {
			return current, parent
		}
		if key < current.data {
			parent = current
			current = current.left
		} else {
			parent = current
			current = current.right
		}
	}

	return current, parent
}

func (b *BinarySearchTree) getMinNodeWithParent(root *BinaryTree) (*BinaryTree, *BinaryTree) {
	var parent *BinaryTree
	// 最小はひたすらに左
	iterator := root
	for iterator.left != nil {
		parent = iterator
		iterator = iterator.left
	}
	return iterator, parent
}

// 配列文字列にして返す
func (b *BinarySearchTree) BstToArr() string {
	// 前順走査
	result := []int32{}
	preorderHelper(b.root, &result)
	return fmt.Sprint(result)
}

func preorderHelper(node *BinaryTree, result *[]int32) {
	if node == nil {
		return
	}
	*result = append(*result, node.data)
	preorderHelper(node.left, result)
	preorderHelper(node.right, result)
}

// ===============Util=============================
// 配列をソートする処理
func mergeSort(arr []int32) []int32 {
	if len(arr) == 1 {
		return arr
	}
	// 半分にしていく処理
	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	i, j := 0, 0
	mergedArr := make([]int32, 0)
	// 小さいもの順に勝ち抜けする処理
	for len(left) > i && len(right) > j {
		if left[i] <= right[j] {
			mergedArr = append(mergedArr, left[i])
			i++
		} else {
			mergedArr = append(mergedArr, right[j])
			j++
		}
	}

	mergedArr = append(mergedArr, left[i:]...)
	mergedArr = append(mergedArr, right[j:]...)

	return mergedArr
}

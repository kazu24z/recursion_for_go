package main

import "fmt"

func main() {
	arr := []int32{3, 6, 1, 4, 9, 2, 8, 7, 5}

	bst := NewBinarySearchTree(arr)
	bst.Insert(11)
	fmt.Println(bst.root.data)
	fmt.Println(bst.root.right.data)
	fmt.Println(bst.root.right.right.data)
	fmt.Println(bst.root.right.right.right.data)
	fmt.Println(bst.root.right.right.right.right.data)

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
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSortArr(arr[:mid])  // 0からmidの1つ前までの配列を渡してる
	right := mergeSortArr(arr[mid:]) // midから末尾までを渡してる

	// マージ処理
	result := make([]int32, 0, len(arr))
	i, j := 0, 0 //カウンター

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)  // iは探索の途中状態を持ってるのでそこから
	result = append(result, right[j:]...) // jも

	return result
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

package main

// "fmt"
// "strings"

func minWindowArrK(intArr []int32, k int32) []int32 {
	deque := &Deque{items: []int32{}} // 空のDeque（インデックスを保持）
	result := make([]int32, 0)

	// 配列を1つずつ走査
	for i := int32(0); i < int32(len(intArr)); i++ {
		// 1. 範囲外のインデックスをDequeの先頭から削除
		if !deque.IsEmpty() && deque.Front() < i-k+1 {
			deque.PopFront()
		}

		// 2. Dequeの末尾から、今の値より大きい要素のインデックスを削除
		//    (これらは今後最小値になることはない)
		for !deque.IsEmpty() && intArr[deque.Back()] >= intArr[i] {
			deque.PopBack()
		}

		// 3. 現在のインデックスを追加
		deque.PushBack(i)

		// 4. ウィンドウが完成したら結果に追加
		if i >= k-1 {
			result = append(result, intArr[deque.Front()])
		}
	}

	return result // ← これが抜けてた
}

// Dequeとウインドウを分けて考えるの重要！！
// Dequeはあくまで最小値候補のインデックスを管理
// ウインドウをずらす→それに伴ってDequeの状態を更新するイメージ

type Deque struct {
	items []int32 // インデックスを保持
}

func (d *Deque) PushBack(val int32) {
	d.items = append(d.items, val)
}

func (d *Deque) PopBack() {
	d.items = d.items[:len(d.items)-1]
}

func (d *Deque) PopFront() {
	d.items = d.items[1:]
}

func (d *Deque) Front() int32 {
	return d.items[0]
}

func (d *Deque) Back() int32 {
	return d.items[len(d.items)-1]
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

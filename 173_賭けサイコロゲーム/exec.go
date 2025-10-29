package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println(diceStreakGamble([]int32{1, 2, 3}, []int32{3, 4, 2}, []int32{4, 2, 4}, []int32{6, 16, 4}))

	gain, name, source := calcGain([]int32{1, 2, 3}, "test")
	fmt.Println(name+":", gain, source)
}

func diceStreakGamble(player1 []int32, player2 []int32, player3 []int32, player4 []int32) string {
	type in struct {
		arr  []int32
		name string
	}
	type out struct {
		gain   int32
		name   string
		source string
	}

	players := []in{
		{player1, "Player 1"},
		{player2, "Player 2"},
		{player3, "Player 3"},
		{player4, "Player 4"},
	}

	var (
		wg      sync.WaitGroup
		results = make([]out, len(players)) // インデックスで格納（ロック不要）
	)

	wg.Add(len(players))
	for i, p := range players {
		i, p := i, p // ループ変数を閉じ込める
		go func() {
			defer wg.Done()
			gain, name, source := calcGain(p.arr, p.name)
			results[i] = out{gain: gain, name: name, source: source}
		}()
	}
	wg.Wait()

	w := results[0]
	for _, r := range results[1:] {
		if r.gain > w.gain {
			w = r
		}
	}

	return fmt.Sprintf("Winner: %s won $%d by rolling %v", w.name, w.gain, w.source)
}

// 1人分
func calcGain(arr []int32, player string) (int32, string, string) {
	total := int32(0)
	stack := &SinglyLinkedListNode{}
	rollings := make([]int32, 0)
	for i := 0; i < len(arr); i++ {
		if stack.head == nil || stack.head.data <= arr[i] {
			stack.Push(arr[i])
			total += 4
			// 該当サイコロ目を保存するリストに入れる
			rollings = append(rollings, arr[i])
		} else {
			if stack.head != nil {
				total = 0
				stack.head = nil
				stack.Push(arr[i])
				total += 4
				rollings = rollings[:0]
				rollings = append(rollings, arr[i])
			}
		}
	}
	// カンマ区切りの文字列を作成
	strs := make([]string, len(rollings))
	for i, v := range rollings {
		strs[i] = strconv.Itoa(int(v))
	}
	rollingsStr := "[" + strings.Join(strs, ",") + "]"

	return total, player, rollingsStr
}

type Node struct {
	data int32
	next *Node
}

type SinglyLinkedListNode struct {
	head *Node
}

// headの値を返す関数
func (s *SinglyLinkedListNode) Peak() int32 {
	return s.head.data
}

// StackのPushとPop処理
func (s *SinglyLinkedListNode) Push(data int32) {
	// headに追加していく
	newNode := &Node{data: data, next: s.head}
	s.head = newNode
}

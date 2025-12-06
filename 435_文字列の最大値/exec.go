package main

import (
	"fmt"
)

func main() {
	r1 := maxByCriteria(compareLength, []string{"apple", "yumberry", "grape", "banana", "mandarin"})

	r2 := maxByCriteria(compareLength, []string{"zoomzoom", "choochoo", "beepbeep", "ahhhahhh"})

	r3 := maxByCriteria(compareAsciiTotal, []string{"apple", "yumberry", "grape", "banana", "mandarin"})

	r4 := maxByCriteria(compareAsciiTotal, []string{"zoom", "choochoo", "beepbeep", "ahhhahhh"})

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
}

type CompareFunc func(string, string) bool

func maxByCriteria(callback CompareFunc, arr []string) string {
	// callbackを基準として、勝ち残りでチェックする
	// 1回目はarr[0], arr[1]で勝負
	// max = 勝ったほう
	// ２回目以降はarr[i]とmaxで勝負
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if !callback(max, arr[i]) {
			max = arr[i]
		}
	}
	return max
}

func compareLength(s1 string, s2 string) bool {
	// 問題は「以上」って言ってんのに、>にしないと同じ文字数のときに更新されない
	if len(s1) > len(s2) {
		return true
	} else {
		return false
	}
}

func compareAsciiTotal(s1 string, s2 string) bool {
	f := func(s string) int {
		r := []rune(s)
		sum := 0
		for i := 0; i < len(s); i++ {
			sum += int(r[i])
		}
		return sum
	}

	sum1 := f(s1)
	sum2 := f(s2)
	if sum1 > sum2 {
		return true
	} else {
		return false
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// テストコード例
	result1 := getNumber("DD")
	fmt.Println("答え：", result1)

	result2 := getNumber("III")
	fmt.Println(result2)

	result3 := getNumber("IDID")
	fmt.Println(result3)
}

func getNumber(code string) string {
	// DでもIでもない場合
	if !isDIOnly(code) {
		fmt.Println("isDIOnly → false")
		return "0"
	}

	codeNum := len(code)

	// codeの文字数が8以上の場合
	if codeNum > 8 {
		return "0"
	}

	// スタック（配列）
	stack := make([]int, 0)
	stack = append(stack, 1) // 1を入れた状態でスタート

	var builder strings.Builder

	// スタックに入れていく値
	for i := 0; i < codeNum; i++ {
		// その回の文字がDならスタックに値を追加
		if code[i] == 'D' {
			stack = append(stack, i+2)
		} else if code[i] == 'I' {
			// IならスタックをPopし、result配列にFIFOで入れる
			for len(stack) > 0 {
				var poppedValue int
				poppedValue, stack = popStack(stack)
				builder.WriteString(strconv.Itoa(poppedValue)) // 取り出した数値を文字列化してbuilderに渡す
			}
			stack = append(stack, i+2)
		}
	}

	// 最終的に残っているスタックの処理
	for len(stack) > 0 {
		var poppedValue int
		poppedValue, stack = popStack(stack)
		builder.WriteString(strconv.Itoa(poppedValue)) // 取り出した数値を文字列化してbuilderに渡す
	}

	return builder.String() // 最終的な文字列を返す
}

func popStack(stack []int) (int, []int) {
	poppedValue := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return poppedValue, stack
}

func isDIOnly(s string) bool {
	// 1. 文字列が空の場合は、構成要素がないため false とする
	if len(s) == 0 {
		fmt.Println("isDIOnly > len(s) == 0")
		return false
	}

	// 2. 文字列をルーンごとに反復処理する
	for _, r := range s {
		// 各ルーンが 'D' または 'I' のどちらでもない場合、即座に false を返す
		if r != 'D' && r != 'I' {
			fmt.Printf("%d", r)
			return false
		}
	}

	// 3. 全ての文字がチェックをパスしたら true を返す
	return true
}

// 文字1つに対してスタックに1から値を入れていく。
// "I"が来たらそれまでのスタックの数値をpopして並びの数字とする
// "I"のときの数字をスタックの底に入れて次の処理へ

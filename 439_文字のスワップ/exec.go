package main

import (
	"fmt"
)

func main() {
	fmt.Println(swapCase([]byte{'l', 'A', 'm', 'b', 'D', 'A'}))
	fmt.Println(swapCase([]byte{'a', 'E', 's', 'P', 'A'}))
	fmt.Println(swapCase([]byte{'f', 'U', 'N', 'c', 'T', 'I', 'o', 'n'}))

}

func swapCase(charList []byte) []byte {
	// 関数を完成させてください
	myMap := func(c byte) byte {
		if 'a' <= c && c <= 'z' {
			return c - 'a' + 'A'
		} else if 'A' <= c && c <= 'Z' {
			return c - 'A' + 'a'
		}
		return c
	}

	var result []byte
	for i := 0; i < len(charList); i++ {
		result = append(result, myMap(charList[i]))
	}
	return result
}

// バイト列それぞれに対して、大小入れ替える処理を呼ぶ

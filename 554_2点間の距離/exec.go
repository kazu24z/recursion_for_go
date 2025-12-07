package main

import (
	"strconv"
)

func calcDistanceList(pointsList []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(pointsList); i++ {
		result = append(result, calcDistance(pointsList[i]))
	}
	return result
}

// 文字列をハイフン区切りで整数に分解
// A,B地点の差を計算（絶対値にする）
// 返す
func calcDistance(point string) string {
	delimiterIndex := 0
	for i := 0; i < len(point); i++ {
		if point[i] == '-' {
			delimiterIndex = i
		}
	}
	x, err := strconv.Atoi((point[:delimiterIndex]))
	if err != nil {
		panic("数値にできません。")
	}
	y, err := strconv.Atoi((point[delimiterIndex+1:]))
	if err != nil {
		panic("数値にできません。")
	}

	abs := 0
	if x > y {
		abs = x - y
	} else {
		abs = y - x
	}
	return strconv.Itoa(abs)
}

package main

import "fmt"

func main() {
	arr := []int32{45, 65, 20, 3, 4, 5, 66, 19, 23, 3, 1}
	result1 := shipmentVolumePackages(arr)
	fmt.Println("結果：", result1)
}

func shipmentVolumePackages(packages []int32) int32 {
	result := int32(0)
	for len(packages) > 1 {
		// 最小値と第二最小値を特定
		min1Index, min2Index := findMins(packages)
		// 最小値を足した値を格納する
		pack := packages[min1Index] + packages[min2Index]
		// 結果に加算
		result += pack

		// 使用した配列要素を削除
		// 大きいインデックスから削除しないと、インデックスがずれる
		if min1Index > min2Index {
			packages = deleteElement(packages, min1Index)
			packages = deleteElement(packages, min2Index)
		} else {
			packages = deleteElement(packages, min2Index)
			packages = deleteElement(packages, min1Index)
		}

		// packages配列の末尾に追加
		packages = append(packages, pack)
	}

	return result
}

// 配列から最小値と第二最小値のindexを見つける処理
func findMins(arr []int32) (int32, int32) {

	// 要素が1つ以下なら
	if len(arr) == 1 {
		return 0, -1 // インデックス0と、2番目がない場合は-1
	} else if len(arr) == 0 {
		return -1, -1 // 両方とも存在しない
	}

	// 線形探索
	// インデックスを追跡
	var minIdx1 int32
	var minIdx2 int32
	if arr[0] > arr[1] {
		minIdx1 = 1
		minIdx2 = 0
	} else {
		minIdx1 = 0
		minIdx2 = 1
	}

	for i := 2; i < len(arr); i++ {
		// arr[i]がmin1より小さい場合：minIdx1 = i, minIdx2 = 前のminIdx1
		if arr[i] <= arr[minIdx1] {
			minIdx2 = minIdx1
			minIdx1 = int32(i)
		} else if arr[i] < arr[minIdx2] {
			// arr[i]がmin1より大きく、min2より小さい場合: minIdx2 = i
			minIdx2 = int32(i)
		}
	}

	return minIdx1, minIdx2
}

func deleteElement(arr []int32, index int32) []int32 {
	lastIndex := len(arr) - 1
	arr[index] = arr[lastIndex]
	arr = arr[:lastIndex]
	return arr
}

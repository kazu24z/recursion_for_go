package main

func dailyStockPrice(stocks []int32) []int32 {
	length := int32(len(stocks))

	// 結果配列
	result := make([]int32, 0)

	// 配列の要素ごとに回す処理
	for i := int32(0); i < length; i++ {

		if i == length-1 {
			result = append(result, 0)
			break
		}
		for j := i + 1; j < length; j++ {
			// stocks[i] < stocks[j]になったときのjの値が上昇まで待つ日
			if stocks[i] < stocks[j] {
				result = append(result, j-i)
				break
			}

			// 最後のループになってもstocks[i] < stocks[j]にならない場合
			if stocks[i] >= stocks[j] && j == length-1 {
				result = append(result, 0)
				break
			}
		}
	}

	return result
}

// stack使わなかった...

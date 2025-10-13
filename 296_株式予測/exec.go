package main

func dailyStockPrice(stocks []int32) []int32 {
	length := int32(len(stocks))

	// 結果配列を全部0で初期化
	result := make([]int32, length)

	// スタック（インデックスを入れる）
	stack := []int32{}

	// 配列を右から左に走査
	for i := length - 1; i >= 0; i-- {
		// スタックトップが現在の株価以下なら全部捨てる
		for len(stack) > 0 && stocks[stack[len(stack)-1]] <= stocks[i] {
			stack = stack[:len(stack)-1] // pop
		}

		// スタックに何か残ってたら、それが答え
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}
		// スタックが空なら result[i] = 0 のまま

		// 今日をスタックに追加
		stack = append(stack, i)
	}

	return result
}

package main

func stockSpan(stocks []int32) []int32 {
	stack := make([]int32, 0)
	results := make([]int32, 0)

	for i := int32(0); i < int32(len(stocks)); i++ {
		// 現在の株価以下のものをスタックから除去
		for len(stack) > 0 && stocks[stack[len(stack)-1]] <= stocks[i] {
			stack = stack[:len(stack)-1]
		}

		// スパンを計算
		var span int32
		if len(stack) == 0 {
			span = i + 1 // 全ての過去の日が含まれる
		} else {
			span = i - stack[len(stack)-1] // 最後に大きかった日との差
		}

		results = append(results, span)
		stack = append(stack, i) // 現在のインデックスをスタックに追加
	}

	return results
}

// 配列のindexをスタックに入れていく
// 各回で勝負してそのindexの差を別の配列に入れていく
//// これが結果配列（先頭分から値が入っていく）

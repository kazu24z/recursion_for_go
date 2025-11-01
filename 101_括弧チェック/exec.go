package main

func isParenthesesValid(parentheses string) bool {
	runes := []rune(parentheses) // 入力文字列をint32の数値の配列に変換
	/** 自分用メモ
	Go言語 という文字列 stringがあり、これの人間が見た文字数（この場合は4文字）を出したい場合
	→ 各文字をUnicodeコードポイント（Unicodeに対応した数値）に変換する
	→ これがrune。実体はint32だが、rune型の値は、その数値に対応する文字という意味を付与する
	→ 文字列をrune配列 []rune() にキャストすると文字ごとに配列要素ができる。なのでこれをlen()したら文字数がわかる
	*/

	// 奇数文字数はその時点でfalse
	if len(runes)%2 != 0 {
		return false
	}
	// 開く側：), ], } をkey
	// 閉じる側：  (, [, { を値にしたmap
	// bracket := make(map[rune]rune)
	// bracket[')'] = '('
	// bracket[']'] = '['
	// bracket['}'] = '{'

	// リテラルで一撃で初期化できる
	bracket := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, 0)

	// 入力文字列を処理
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '(', '{', '[':
			// スタック(リスト末尾)に値を格納する
			stack = append(stack, runes[i])
		case ')', '}', ']':
			// スタックの先頭がペア文字の場合
			// if len(stack) > 0 && bracket[runes[i]] == stack[len(stack)-1] {
			// 	// スタックからPopする
			// 	stack = stack[:len(stack)-1]
			// } else {
			// 	// ペアが成り立たない時点で終了
			// 	return false
			// }

			// 条件を逆転させるとコードが簡潔になる
			// NG条件→即return false, それ以外はPop
			if len(stack) == 0 || bracket[runes[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	// ループ処理後にstackに値が残っていてもfalse "{[("パターン
	if len(stack) > 0 {
		return false
	}

	return true
}

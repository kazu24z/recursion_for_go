## 数式の解析（応用） hard
問題 295 では文字列で表現された数式を評価して整数の結果を返すことができました。括弧 () が含まれた文字列 expression が与えられるので、その式を解析して整数の結果を返す、expressionParenthesisParser という関数を作成してください。四則演算の際、割り算の結果は全て切り捨てを行ってください。


関数の入出力例
```
入力のデータ型： string expression

出力のデータ型： long

expressionParenthesisParser("2+4*6") --> 26

expressionParenthesisParser("2*3+4") --> 10

expressionParenthesisParser("3-3+3") --> 3

expressionParenthesisParser("2+2+2") --> 6

expressionParenthesisParser("1-1-1") --> -1

expressionParenthesisParser("3*3/3*3*3") --> 27

expressionParenthesisParser("42") --> 42

expressionParenthesisParser("7*3622*636*2910*183+343/2926/1026") --> 8587122934320

expressionParenthesisParser("(2*3)+(1+2)") --> 9

expressionParenthesisParser("4/(486-484)") --> 2

expressionParenthesisParser("(1+(2+3+4)-3)+(9+5)") --> 21

expressionParenthesisParser("(100+300)*5+(20-10)/10") --> 2001

expressionParenthesisParser("(100+200)/3*100+1000/10") --> 10100

```

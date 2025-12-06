## 文字列の最大値 medium
maxByCriteria 関数は、文字列の配列と特定の比較基準を適用するコールバック関数を受け取り、その基準に従って配列内の最大の文字列を返します。提供されるコールバック関数の例は、以下の通りです。

bool compareLength(string s1, string s2)
文字列 s1 と s2 を受け取り、s1 の長さが s2 の長さ以上の場合に true を返し、それ以外の時に false を返す関数。

bool compareAsciiTotal(string s1, string s2)
文字列 s1 と s2 を受け取り、s1 の ASCII 値の合計が s2 の合計以上の場合に true を返し、それ以外の時に false を返す関数。

maxByCriteria 関数

入力：
（Function）比較関数の参照
（array）文字列の配列
出力：
（str 型）配列内で最大と判断される文字列
関数の入出力例
```
maxByCriteria(compareLength, ["apple", "yumberry", "grape", "banana","mandarin"]) --> mandarin

maxByCriteria(compareLength, ["zoomzoom", "choochoo", "beepbeep", "ahhhahhh"]) --> ahhhahhh

maxByCriteria(compareAsciiTotal, ["apple", "yumberry", "grape", "banana","mandarin"]) --> yumberry

maxByCriteria(compareAsciiTotal, ["zoom", "choochoo", "beepbeep", "ahhhahhh"]) --> choochoo

```

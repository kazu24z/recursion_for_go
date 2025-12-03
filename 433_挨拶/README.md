## 挨拶 easy
Knott は、スマートスピーカーの開発者です。彼はスピーカーがユーザーにカスタマイズされた挨拶をする機能を実装しています。この機能では、時間帯に応じた挨拶と、ユーザーが指定したテキスト変換操作を組み合わせて挨拶します。具体的には、時間 time、テキスト text 、およびテキストを変換する関数 f を受け取り、カスタマイズされた挨拶メッセージを生成する greet 関数を作成する必要があります。

挨拶のルールは以下の通りです。

time が 0 以上、12 より小さい場合は "Good Morning"
time が 12 以上、18 より小さい場合は "Good Afternoon"
それ以外の場合は "Good Evening"
さらに、text に対して以下のテキスト変換関数 f のいずれかを適用します。

string loud(string text)
text を文字列として受け取り、すべて大文字で返す関数。

string quiet(string text)
text を文字列として受け取り、すべて小文字で返す関数。

string reverse(string text)
text を文字列として受け取り、文字の順序を逆転して返す関数。

string repeat(string text)
text を文字列として受け取り、文字列を 2 回繰り返す関数。

最終的な挨拶メッセージは、選択された時間帯に応じた挨拶と、変換されたテキストを組み合わせたものになります。

greet 関数

入力：
（int 型）時間
（str 型）名前
（Function）テキストを処理する関数への参照
出力：
（str 型）指定された時間に応じた挨拶と、テキストに適用された関数の結果

関数の入出力例
```
greet(1, "John", loud) --> Good Morning JOHN

greet(2, "John", quiet) --> Good Morning john

greet(13, "John", reverse) --> Good Afternoon nhoJ

greet(19, "John", repeat) --> Good Evening John John

greet(13, "Leslie Emmanuel Beadon", loud) --> Good Afternoon LESLIE EMMANUEL BEADON

greet(19, "Leslie Emmanuel Beadon", quiet) --> Good Evening leslie emmanuel beadon

greet(5, "Leslie Emmanuel Beadon", reverse) --> Good Morning nodaeB leunammE eilseL

greet(1, "Leslie Emmanuel Beadon", repeat) --> Good Morning Leslie Emmanuel Beadon Leslie Emmanuel Beadon

```

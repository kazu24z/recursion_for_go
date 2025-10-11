## 片方向リストn倍 medium
片方向リストの先頭 head、自然数 n が与えられるので、n 倍した連結リストを返す、reproduceByN という関数を作成してください。

関数の入出力例
```
入力のデータ型： SinglyLinkedListNode<integer> head, integer n

出力のデータ型： SinglyLinkedListNode<integer>

reproduceByN(singlyLinkedList([2,10,34,45,67,356]),3) --> 2➡10➡34➡45➡67➡356➡2➡10➡34➡45➡67➡356➡2➡10➡34➡45➡67➡356➡END

reproduceByN(singlyLinkedList([20,-5,34,45,67,356]),4) --> 20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡END

reproduceByN(singlyLinkedList([20,-5,34,45,67,356]),6) --> 20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡20➡-5➡34➡45➡67➡356➡END

reproduceByN(singlyLinkedList([-8,14,34,45,67,356]),1) --> -8➡14➡34➡45➡67➡356➡END
```

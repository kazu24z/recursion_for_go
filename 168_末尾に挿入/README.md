## 末尾に挿入 easy

連結リストの先頭 head と整数 data が与えられるので、連結リストの末尾にデータを挿入した新しい連結リストを返す、insertAtTail という関数を作成してください。

関数の入出力例

```

入力のデータ型： SinglyLinkedListNode<integer> head, integer data

出力のデータ型： SinglyLinkedListNode<integer>

insertAtTail(singlyLinkedList([3,3,2,10,34,45,67,356]),367) --> 3➡3➡2➡10➡34➡45➡67➡356➡367➡END

insertAtTail(singlyLinkedList([3,3,2,10,34,45,67,356]),4) --> 3➡3➡2➡10➡34➡45➡67➡356➡4➡END

insertAtTail(singlyLinkedList([3,8]),4) --> 3➡8➡4➡END

insertAtTail(singlyLinkedList([3,32,2,10,34,45,67,356]),46) --> 3➡32➡2➡10➡34➡45➡67➡356➡46➡END

insertAtTail(singlyLinkedList([20,-5,34,45,67,356,34,687,31,-34,5]),75) --> 20➡-5➡34➡45➡67➡356➡34➡687➡31➡-34➡5➡75➡END

insertAtTail(singlyLinkedList([20,-5,34,45,67,356,34,687,31,-34,5]),96) --> 20➡-5➡34➡45➡67➡356➡34➡687➡31➡-34➡5➡96➡END
```

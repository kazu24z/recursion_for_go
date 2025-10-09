## 特定位置に挿入 easy
ソート済みの連結リストの先頭 head と、データ data が与えられるので、ノードを正しい位置に挿入した連結リストを返す、insertNodeInSorted という関数を作成してください。
関数の入出力例
```
入力のデータ型： SinglyLinkedListNode<integer> head, integer data

出力のデータ型： SinglyLinkedListNode<integer>

insertNodeInSorted(singlyLinkedList([2,10,34,45,67,356]),3) --> 2➡3➡10➡34➡45➡67➡356➡END

insertNodeInSorted(singlyLinkedList([2,10,34,45,67,356]),358) --> 2➡10➡34➡45➡67➡356➡358➡END

insertNodeInSorted(singlyLinkedList([2,10,34,45,67,356]),-5) --> -5➡2➡10➡34➡45➡67➡356➡END

insertNodeInSorted(singlyLinkedList([4,23,53,56,134,645]),34) --> 4➡23➡34➡53➡56➡134➡645➡END

```

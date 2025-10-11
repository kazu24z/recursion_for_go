## 連結リストの交わる点 medium
連結リストの先頭 headA と headB がそれぞれ与えられるので、お互いの連結リストが交わるノードの値を返す、findMergeNode という関数を作成してください。headA と headB は異なるものとし、null ではありません。また全ての整数は正とし、A と B の交わりが存在しない場合は -1 を返してください。

関数の入出力例
```
入力のデータ型： SinglyLinkedListNode<integer> headA, SinglyLinkedListNode<integer> headB

出力のデータ型： integer

findMergeNode(singlyLinkedList([2,10,34,45,67,356]),singlyLinkedList([20,5,34,45,67,356])) --> 34

findMergeNode(singlyLinkedList([34,657,24,36,75,867,345,75,345,64,75]),singlyLinkedList([13,4,6,3,345,64,75])) --> 345

findMergeNode(singlyLinkedList([34,657,24,36,75,867,345,75,345,64,75]),singlyLinkedList([13,4,6,3,345,64,75,34])) --> -1

```

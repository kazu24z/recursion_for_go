## ソート済連結リストの合併 medium
ソート済みの連結リストの先頭 head1 と head2 が与えられるので、それを合体させ、新しい連結リストを返す、mergeTwoSortedLinkedLists という関数を作成してください。

連結リストの先頭 head と整数 data が与えられるので、リストの先頭にデータを挿入した新しい連結リストを返す、insertAtHead という関数を作成してください


関数の入出力例
```

入力のデータ型： SinglyLinkedListNode<integer> head1, SinglyLinkedListNode<integer> head2

出力のデータ型： SinglyLinkedListNode<integer>

mergeTwoSortedLinkedLists(singlyLinkedList([1,2,3]),singlyLinkedList([1,4,5])) --> 1➡1➡2➡3➡4➡5➡END

mergeTwoSortedLinkedLists(singlyLinkedList([1,2,3,10,30,45]),singlyLinkedList([1,4,30,50,80])) --> 1➡1➡2➡3➡4➡10➡30➡30➡45➡50➡80➡END

mergeTwoSortedLinkedLists(singlyLinkedList([2,3,5]),singlyLinkedList([5,8,12])) --> 2➡3➡5➡5➡8➡12➡END

mergeTwoSortedLinkedLists(singlyLinkedList([1,1,2,3,3,5]),singlyLinkedList([2,2,3,3,5,5,10,10])) --> 1➡1➡2➡2➡2➡3➡3➡3➡3➡5➡5➡5➡10➡10➡END
```

## Dequeクラス easy
両端キューとは、先頭または末尾で要素を追加、削除できるキューのことを指します。両端キューは 4 つの基本操作 enqueueFront, enqueueBack, dequeueFront, dequeueBack を持ちます。

両端キューに含まれる要素は、以下のように双方向リストに基づいた Item クラスによって管理することができます。

int data: 要素の値
Item prev: 1 つ前のノード。デフォルトでは null に設定してください。
Item next: 1 つ先のノード。デフォルトでは null に設定してください。
両端キューでは、先頭と末尾を定義することによって、要素の追加や削除をスムーズに行うことができます。以下に従って、Deque クラスを作成してください。

Item head: 先頭のノード。デフォルトでは null に設定してください。
Item tail: 末尾のノード。デフォルトでは null に設定してください。
Integer peekFront(): キューの先頭の Item の値を返します。
Integer peekBack(): キューの末尾の Item の値を返します。
void enqueueFront(int data): 先頭に Item を追加します。
void enqueueBack(int data): 末尾に Item を追加します。
Integer dequeueFront(): 先頭から Item を削除し、その値を返します。両端キューに Item が存在しない場合、null を返してください。
Integer dequeueBack(): 末尾から Item を削除し、その値を返します。両端キューに Item が存在しない場合、null を返してください。

関数の入出力例
```

q = new Deque()

q.enqueueBack(4)

q.enqueueBack(50)

q.peekFront() --> 4

q.peekBack() --> 50

q.enqueueFront(36)

q.enqueueFront(96)

q.peekFront() --> 96

q.peekBack() --> 50

q.dequeueBack() --> 50

q.dequeueBack() --> 4

q.peekFront() --> 96

q.peekBack() --> 36

q.enqueueFront(20)

q.dequeueBack() --> 36
```

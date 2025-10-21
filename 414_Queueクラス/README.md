## Queueクラス easy
キューとはコンピュータで頻繁に使用される特殊なデータ構造です。キューは最初に入ってきた要素が最初に実行されるという構造で、2 つの基本操作 enqueue と dequeue を持ちます。

キューに含まれる要素は、以下のように、片方向リストに基づいた Node クラスによって管理することができます。

- int data: 要素の値
- Node next: 1 つ先のノード。デフォルトでは null に設定してください。
以下のように、先頭と末尾を定義することによって、キューをデータとして表現することができます。
以下に従って、Queue クラスを作成してください。

- Node head: 先頭のノード。デフォルトでは null に設定してください。
- Node tail: 末尾のノード。デフォルトでは null に設定してください。
- Integer peekFront(): キューの先頭の Node の値を返します。
- Integer peekBack(): キューの末尾の Node の値を返します。
- void enqueue(int data): キューの末尾に Node を挿入します。
- Integer dequeue(): キューの先頭の Node を取り除き、その値を返します。

テストケース
```
q = new Queue()

q.enqueue(4)

q.peekFront() --> 4

q.peekBack() --> 4

q.enqueue(50)

q.peekFront() --> 4

q.peekBack() --> 50

q.enqueue(64)

q.peekFront() --> 4

q.peekBack() --> 64

q.dequeue() --> 4

```

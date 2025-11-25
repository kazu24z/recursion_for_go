## 対称的な二分木 medium
二分木 root が与えられるので、それが左右対称かどうかを確かめる、symmetricTree という関数を作成してください。

関数の入出力例
```
入力のデータ型： binaryTree<integer> root

出力のデータ型： bool

symmetricTree( toBinaryTree([10,25,25,33,45,45,33]) )--> true 

symmetricTree( toBinaryTree([10,25,25,33,45,45,32]) )--> false 

symmetricTree( toBinaryTree([1,2,2,3,4,4,3]) )--> true 

symmetricTree( toBinaryTree([3,6,6,9,12,11,9]) )--> false 

symmetricTree( toBinaryTree([]) )--> true 

symmetricTree( toBinaryTree([1,9,9,15,19,19,15]) )--> true 

symmetricTree( toBinaryTree([1,2,2,null,3,null,3]) )--> false 

symmetricTree( toBinaryTree([3,6,6,9,12,12,9]) )--> true 

symmetricTree( toBinaryTree([3,6,7,9,12,12,9]) )--> false 

symmetricTree( toBinaryTree([1,2,2,2,null,2]) )--> false 

symmetricTree( toBinaryTree([1,2,3]) )--> false 

```

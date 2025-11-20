## 祖父母ノード medium
自然数で構成される二分木 root が与えられるので、祖父母ノードが偶数値を持つノードの合計を返す、totalEvenGrandparent という関数を作成してください。ここでの祖父母ノードは親ノードの親ノードを指します。もし、偶数値を持つ祖父母ノードが存在しない場合、0 を返してください。

関数の入出力例
```
入力のデータ型： binaryTree<integer> root

出力のデータ型： integer

totalEvenGrandparent( toBinaryTree([8,9,11]) )--> 0 

totalEvenGrandparent( toBinaryTree([8,9,11,3,null,null,2]) )--> 5 

totalEvenGrandparent( toBinaryTree([57,33,77,9,40,61,92,7,14,35,46,59,63,78,96,null,null,null,23,null,37,null,47,null,null,null,76,null,84,null,99]) )--> 267 

totalEvenGrandparent( toBinaryTree([41,15,70,8,28,55,78,4,10,21,34,47,63,74,83,2,6,9,14,16,25,33,35,46,50,56,65,72,75,81,90,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,40,null,null,null,null,null,null,null,69,null,null,null,null,null,null,null,96]) )--> 765 

totalEvenGrandparent( toBinaryTree([52,33,72,17,39,63,85,10,25,35,44,58,69,80,98,5,12,18,27,34,36,40,48,53,59,66,70,73,82,89,99,null,null,null,16,null,22,null,28,null,null,null,37,null,43,null,49,null,null,null,60,null,68,null,71,null,79,null,83,null,94,null,99]) )--> 1032 

```

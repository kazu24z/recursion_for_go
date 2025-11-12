## 二分木の最大の深さ medium
二分木 root が与えられるので、最大の深さを返す、maximumDepth という関数を作成してください。ここでの最大の深さとは、根ノードから最も遠い葉ノードまでの最長経路に沿ったノードの数を指します。

関数の入出力例
```
入力のデータ型： binaryTree<integer> root

出力のデータ型： integer

maximumDepth( toBinaryTree([0]) )--> 0 

maximumDepth( toBinaryTree([0,1,null]) )--> 1 

maximumDepth( toBinaryTree([0,-10,5,null,-3,null,9]) )--> 2 

maximumDepth( toBinaryTree([5,2,18,-4,3]) )--> 2 

maximumDepth( toBinaryTree([27,14,35,10,19,31,42]) )--> 2 

maximumDepth( toBinaryTree([30,15,60,7,22,45,75,null,null,17,27]) )--> 3 

maximumDepth( toBinaryTree([90,50,150,20,75,95,175,5,25,66,80,92,111,166,200]) )--> 3 

maximumDepth( toBinaryTree([50,17,76,9,23,54,null,null,14,19,null,null,72]) )--> 3 

maximumDepth( toBinaryTree([16,14,10,8,7,9,3,2,4,1]) )--> 3 

maximumDepth( toBinaryTree([30,15,60,7,22,45,75,1,null,17,27,null,null,null,null,-1]) )--> 4 

```

## 有効な二分木 medium
二分木 root が与えられるので、それが二分探索木（BST）かどうか確かめる、validateBST という関数を作成してください。ただし、有効な BST とは以下の条件を指します。

ノードの左部分木にはノードのキーより小さい値を含みます
ノードの右部分木にはノードのキーより大きい値を含みます
左右の部分木も BST である必要があります

関数の入出力例
```

入力のデータ型： binaryTree<integer> root

出力のデータ型： bool

validateBST( toBinaryTree([0,null,-1]) )--> false 

validateBST( toBinaryTree([4,1,5]) )--> true 

validateBST( toBinaryTree([15,10,20,8,12,16,25]) )--> true 

validateBST( toBinaryTree([30,15,60,7,22,45,75,null,null,17,27]) )--> true 

validateBST( toBinaryTree([5,1,2,3,4]) )--> false 

validateBST( toBinaryTree([5,1,4,null,null,3,6]) )--> false 

validateBST( toBinaryTree([]) )--> true 

validateBST( toBinaryTree([5,3,9,1,6,8]) )--> false 

validateBST( toBinaryTree([1]) )--> true 

validateBST( toBinaryTree([1,2]) )--> false 
```

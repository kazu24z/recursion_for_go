## 完全二分木 medium
二分木が与えられるので、完全二分木（complete binary tree）かどうか判定する、isCompleteBinaryTree という関数を作成してください。完全二分木とは、最下層を除いてすべての深さがノードで満たされ、最下層の葉ノードが可能な限り左に寄せられているような木のことを指します。


関数の入出力例
```
入力のデータ型： binaryTree<integer> root

出力のデータ型： bool

isCompleteBinaryTree( toBinaryTree([0,1,2,3]) )--> true 

isCompleteBinaryTree( toBinaryTree([0,1,2,3,4,5]) )--> true 

isCompleteBinaryTree( toBinaryTree([0,1,2,4,5,null,7]) )--> false 

isCompleteBinaryTree( toBinaryTree([0,1,3,4,null,7,8]) )--> false 

isCompleteBinaryTree( toBinaryTree([0,1,2,3,null,6]) )--> false 

isCompleteBinaryTree( toBinaryTree([97,10,77,32,40,70,32,96,27,10,12,90,73,100,86,null,96]) )--> false 

isCompleteBinaryTree( toBinaryTree([11,7,75,9,59,83,60,46,6,12,26,28,26,91,98,83,7,null,31,66,77,23,null,100,40]) )--> false 

isCompleteBinaryTree( toBinaryTree([38,67,22,52,0,29,43,41,55,97,82,33,85,5,80,3,94,46,0,32,88,59,59,38,64,83,78,47,13,41,89,90,17,36,53,56,1,8,36,92,8,null,78,33,81,86,null,10,6,31,27]) )--> false 
```

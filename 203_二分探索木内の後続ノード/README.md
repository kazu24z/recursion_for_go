## 二分探索木内の後続ノード medium
異なる整数値で構成される二分探索木（BST）の根ノード root と BST 内に存在する整数 key が与えられるので、根ノードが後続ノードである部分木を返す、successor という関数を作成してください。もし、そのようなノードが存在しない場合は null を返してください。ノード N の値を x とした時、後続ノードとは木の中に存在する x よりも大きい最小の値を持つノードのことを指します。


関数の入出力例
```
入力のデータ型： binaryTree<integer> root, integer key

出力のデータ型： binaryTree<integer>

successor( toBinaryTree([0,-10,5,null,-3,null,9]), 5 )--> [9] 

successor( toBinaryTree([5,3,6,2,4,null,7]), 5 )--> [6,null,7] 

successor( toBinaryTree([10,6,12,4,8,null,null,2]), 12 )--> [] 

successor( toBinaryTree([10,6,12,4,8,null,null,2]), 2 )--> [4,2] 

successor( toBinaryTree([5,4,null]), 5 )--> [] 

successor( toBinaryTree([-2,-17,8,-18,-11,3,19,null,null,null,-4,null,null,null,25]), 8 )--> [19,null,25] 

successor( toBinaryTree([3,-3,13,-7,1,6,18,-10,-4,0,2,5,8,15,19]), 6 )--> [8] 

successor( toBinaryTree([3,-3,13,-7,1,6,18,-10,-4,0,2,5,8,15,19]), 3 )--> [5] 

successor( toBinaryTree([1,-5,15,-9,-4,10,17,null,-6,null,0,null,14,16,19]), 10 )--> [14] 

successor( toBinaryTree([0,-10,5,null,-3,null,9]), -3 )--> [0,-10,5,null,-3,null,9] 

```

## ラムダ総和 easy
特定の条件を満たす自然数の総和を計算する関数 summation を作成してください。summation 関数は、与えられたラムダ関数を使用して自然数 n 以下の数値を検証し、条件に合致する数値の総和を計算します。以下の 3 つの条件を判定するラムダ関数を実装し、それらを用いて summation 関数を作成してください。

bool isOdd(int x)
引数 x が奇数かどうかを判定するラムダ関数。奇数の場合は true を、そうでない場合は false を返します。

bool isMultipleOf3Or5(int x)
引数 x が 3 または 5 の倍数かどうかを判定するラムダ関数。倍数の場合は true を、そうでない場合は false を返します。

bool isPrime(int x)
引数 x が素数かどうかを判定するラムダ関数。素数の場合は true を、そうでない場合は false を返します。

summation 関数

入力：
（callable）条件判定
（int 型）上限値となる自然数
出力：
（int 型）総和

関数の入出力例
```
summation(isOdd, 3) --> 4

summation(isOdd, 10) --> 25

summation(isOdd, 25) --> 169

summation(isMultipleOf3Or5, 3) --> 3

summation(isMultipleOf3Or5, 10) --> 33

summation(isMultipleOf3Or5, 100) --> 2418

summation(isPrime, 2) --> 2

summation(isPrime, 10) --> 17

summation(isPrime, 100) --> 1060

```

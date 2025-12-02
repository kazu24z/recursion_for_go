## 動物と人間の年齢 easy
動物の年齢に基づいて医療保険の加入資格を判断する関数 qualifiedForInsurance を作成してください。この関数は、動物の年齢を人間の年齢に換算し、その結果が 60 歳以下であれば true を、60 歳より上であれば false を返す必要があります。動物の年齢を人間の年齢に換算するためのラムダ関数は、以下の 2 つです。

int dogToHuman(int dogAge)
犬の年齢を人間の年齢（int 型）に換算するラムダ関数。換算式は 20 + (犬の年齢 - 2) * 7 です。

int catToHuman(int catAge)
猫の年齢を人間の年齢（int 型）に換算するラムダ関数。換算式は 24 + (猫の年齢 - 2) * 4 です。

qualifiedForInsurance 関数

入力：
（callable）年齢換算
（int 型）動物の年齢
出力：
（bool 型）保険の加入資格があるかどうかを示す

関数の入出力例
```
qualifiedForInsurance(dogToHuman, 7) --> true

qualifiedForInsurance(dogToHuman, 8) --> false

qualifiedForInsurance(catToHuman, 11) --> true

qualifiedForInsurance(catToHuman, 12)) --> false

```

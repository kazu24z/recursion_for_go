## バリデーション easy
Clara はメールアドレスが正しい形式であるかを検証する関数を作成しています。この関数は、指定された検証ルール（validator）に基づいて、メールアドレスをチェックし、結果に応じたメッセージを返す必要があります。以下の 3 つの検証ルールをラムダ関数として実装し、これらを使用して emailValidation 関数を作成してください。

bool doesNotStartWithAt(string email)
email が @ で始まらない場合に true を返し、@ で始まる場合には false を返すラムダ関数。

bool doesNotHaveSpace(string email)
email に空白が含まれていない場合に true を返し、空白が含まれている場合には false を返すラムダ関数。

bool hasUppercaseAndLowercase(string email)
email に大文字と小文字が両方含まれている場合に true を返し、含まれていない場合には false を返すラムダ関数。

各ラムダ関数は特定の条件を満たすかどうかをチェックし、その結果に基づいて emailValidation 関数は適切なメッセージを文字列として返します。これらの関数を使って、指定されたメールアドレスが適切な形式であるかを検証してください。

関数の入出力例
```
emailValidation(doesNotStartWithAt, "@gmail.com") --> Email is not correct.

emailValidation(doesNotStartWithAt, "kkk@gmail.com") --> Email is correct.

emailValidation(doesNotHaveSpace, "Hello world") --> Email is not correct.

emailValidation(doesNotHaveSpace, "Helloworld") --> Email is correct.

emailValidation(hasUppercaseAndLowercase, "hello world") --> Email is not correct.

emailValidation(hasUppercaseAndLowercase, "HELLO WORLD") --> Email is not correct.

emailValidation(hasUppercaseAndLowercase, "Hello world") --> Email is correct.

```

# python

## 第1章 変数とデータ型

    ＞（大なり）
    ＜（小なり）

    文字と数字の相互変換
    ord('文字') と chr(数値) で相互に変換できます。
    https://qiita.com/ell/items/6eb48e934a147898d823


    文字の分割（split）
    https://note.nkmk.me/python-split-rsplit-splitlines-re/

    文字列をリスト関数に入れると１文字ずつ分割
    https://niwakomablog.com/python-list-string/

    ```python
    # python

    name='john'
    type(name)

    # True
    if(isinstance(name,str))

    # False
    if(isinstance(name,int))

    ```

    改行
    print(line, end='')
    https://www.lifewithpython.com/2017/11/python-line-break-handling.html

    フォーマット　０埋め
    https://note.nkmk.me/python-zero-padding/

## 第2章 コレクション+zip関数+emurate関数

    ```python
    # タプル（変更不可）
    t_taple = (1, 2, 3)

    # リスト
    s_list = [1, 2, 3] 
   
    # セット（辞書）
    s_set = {1:10, 2:20, 3:30}
    ```

### 2.5　コレクションの応用

    注意： ()と[]と{}を間違えないように定義しよう

    set（集合型）とlistは違う
    https://aiacademy.jp/media/?p=2941

    リスト　シャッフル
    https://note.nkmk.me/python-random-shuffle/

    ```python
    last_names = ['aaa', 'bbb', 'ccc']

    # 拡張for文
    [print(name) for name in last_names if (isinstance(name, str))]

    # enumerate文（自動でインデックス作ってくれる）
    [print(f'i={i}, name={name}') for i, name in enumerate(last_names) if (isinstance(name, str))]
    
    # zip関数（複数リストを）
    first_list = [1, 2, 3]
    second_list = [10, 20, 30, 40, 50]
    third_list =[100, 200, 300, 400]
    [print(first + second + third) for first, second, third in zip(first_list, second_list, third_list)]
   
    ```

### 第3章 条件分岐

### 第4章 繰り返し

### 第5章 関数

    可変長引数
    https://note.nkmk.me/python-args-kwargs-usage/

    ```python
    def my_sum(*args):
        return sum(args)

    # 10
    print(my_sum(1, 2, 3, 4))
    
    # 36
    print(my_sum(1, 2, 3, 4, 5, 6, 7, 8))
    ```

    引数にリスト、タプル、辞書を展開して渡す
    https://note.nkmk.me/python-argument-expand/

### 第6章 オブジェクト+クラス

クラス

- `__new__` selfを作る（シングルトン用等）
- `__init__` 初期設定
- `__str__` ToString(弱)
- `__repr__` ToString(強)
- `__bool__` 真偽値設定用
- `__eq__` equalsみたいなやつ

        ```python

        class Person:
            # 初期設定(self必須)
            def __init__(self, name, nationality, age):
                self.name = name
                self.nationality = nationality
                self.age = age

            def say_hello(self):
                print(f'私は「{self.name}」です。')

            # JavaのToString同じ
            def __str__(self):
                return f'Name:{self.name}'

            # __str__の詳細版
            def __repr__(self):
                return f'Name:{self.name} Nat:{self.nationality} Age:{self.age}'

            # javaのequalsメソッド
            def __eq__(self, other):
                return True if self.name == other.name \
                            and self.nationality == other.nationality \
                            and self.age == other.age else False


        # student00 = Person()
        student01 = Person(name='first', nationality='jan', age=20)
        student02 = Person(name='second', nationality='jan', age=21)
        student03 = Person(name='third', nationality='usa', age=21)

        # __eq__確認用にstudent00と同じ内容
        student04 = Person(name='first', nationality='jan', age=20)

        print(student01.name, student01.nationality, student01.age)
        student01.say_hello()

        print(str(student01))
        print(repr(student01))

        print(True if student01 == student04 else False)
        ```


![クラス](\img\img_python_class.jpg)

    superクラスの実装

    ```python
    # super style
    class A(object):
        def __init__(self):
            print 'A',

    class B(A):
        def __init__(self):
            print 'B',
            super(B, self).__init__()

    class C(A):
        def __init__(self):
            print 'C',
            super(C, self).__init__()

    class D(B, C):
        def __init__(self):
            print 'D',
            super(D, self).__init__()
    d = D() # D B C A
    ```

    クラス練習用シーザー暗号

    ```python
    class Caesar_cipher:
    def __init__(self, text: str, n: int):
        self.text = text
        self.num = n

    def caesar_encryption(self):
        self.text = "".join([chr(ord(t) + self.num) for t in self.text])
        print(f'暗号化:{self.text}')

    def caesar_composite(self):
        self.text = "".join([chr(ord(t) - self.num) for t in self.text])
        print(f'暗号化:{self.text}')

    def __str__(self):
        return f'登録文字:{self.text} 移行させる数:{self.num}'


    file_name = 'sample1.txt'
    sample_text = 'サンプルテキスト'
    num_of_move_char = 3

    try:
        f = open(file_name, 'r', encoding='UTF-8')
        input_text = f.read()
        f.close()
        print(f'読み込まれたファイルの内容は\n{input_text}\nです。\n')
    except FileNotFoundError:
        print(f'ファイル:{file_name}が無かったので\n{sample_text}\nを代入します。\n')
        input_text = 'サンプルテキスト'


    input01 = Caesar_cipher(input_text, num_of_move_char)
    print(str(input01))
    input01.caesar_encryption()
    input01.caesar_composite()
    ```


第7章 モジュール
7.1　部品を使おう
7.2　組み込み関数
7.3　モジュールの利用
7.4　パッケージの利用
7.5　外部ライブラリの利用
第8章 まだまだ広がるPythonの世界
8.1　Pythonの可能性
8.2　Pythonの基礎を学び終えて
付録A　sukkiri.jpについて

### 付録B　エラー解決・虎の巻

- SyntaxError（構文エラー）
- NameError（定義していない）
- AttributeError（オブジェクトが持っていない属性を使用）
- TypeError（誤ったデータ型どうしの計算）
- ValueError（誤った値）
- IndentationError（インデントミス）
- IndexError（範囲外指定時のエラー）
- KeyError（辞書で登録されていないキーを指定）
- ModuleNotFoundError（インポートミス）
- FileNotFoundError（指定したファイルがない)

### 内包表記

    fizzbuzz問題を内包表記で

    ```python
    def fizz_buzz(n):
    num_list = ['FizzBuzz' if i % 15 == 0
                else 'Fizz' if i % 3 == 0
                else 'Buzz' if i % 5 == 0
                else i
                for i in range(1, n, 1)]
    return num_list


    print(fizz_buzz(100), end="\n")
    ```

### ラムダ式

    ```python
    def calc(base, height):
    return base * height / 2


    print(calc(2, 2))

    num = (lambda base, height: base * height / 2)(5, 10)
    print(num)

    # 小文字のリストを大文字のリストにするlambda処理
    lower_students = ['a', 'b', 'c', 'd']
    upper_students = (lambda student_list: [student.upper() for student in student_list])(lower_students)
    print(upper_students)
    ```

### 可変長引数

### デコレータ

    # 関数内関数

    ```python
    def func1():
        print('func1')


    def func2(f):
        # 関数内関数
        def wrapper():
            print('start')
            f()
            print('end')
        return wrapper


    # func2(func1)()
    func = func2(func1)
    func()
    ```

    # @を使う

    ```python
    def func2(f):
    # 関数内関数
    def wrapper():
        print('start')
        f()
        print('end')

    return wrapper


    # func2でデコレート。メイン処理はfunc1
    @func2
    def func1():
        print('func1')


    func1()
    ```

    # 可変長引数対応デコレータ

    ```python
    def func2(f):
    def wrapper(*args, **kwargs):
        print('start')
        f(*args, **kwargs)
        print('end')
    return wrapper


    @func2
    def func1(*args, **kwargs):
        args_ary = [(str(s)) for s in args]
        kwargs_ary = [(str(s)) for s in kwargs]
        print('func1' + "".join(args_ary) + "".join(kwargs_ary))


    func1('memo1', 'memo2')
    ```

### Type Hints（型ヒント）

    現状は警告のみ（2022/05/03）

    ```python
    def calc_tax_price(p: int, t: float) -> int:
    return int(p * t)


    price: int = 1000
    tax: float = 1.1


    if __name__ == "__main__":
        print(calc_tax_price(price, tax))
    ```



●コラム
Pythonの由来
Python3系を使おう
そのほかの開発環境
円記号とバックスラッシュ
代入演算子の特殊性
Pythonの予約語
複数の単語から作る識別子の命名規則
暗黙の型変換
2つのタイプの関数
f-string
ディクショナリ要素の順序
ディクショナリの合計
コレクションたちの別名
ディクショナリへの変換
1行≠1つの文とならない書き方
チャットボットとAI
タブ文字
文字列の大小比較
範囲指定の条件式
論理演算子の名前の由来
真偽値に評価されない条件式
空ブロックの作り方
三項条件演算子
無限ループを止める方法
「空っぽ」を意味するNone
ディクショナリを用いた可変長引数
関数定義と呼び出しのコーディング
関数さえオブジェクト
コレクション変換関数の正体
「箱」より「名札」に近いPythonの変数
捨てられた不変オブジェクトの行方
破壊的な関数
不変オブジェクトの再利用
ストリーム
車輪の再発明
組み込み関数の正体
外部ライブラリのインストール
MicroPython
R言語
# python

## 第1章 変数とデータ型

＞（大なり）
＜（小なり）

文字と数字の相互変換
ord('文字') と chr(数値) で相互に変換できます。
<https://qiita.com/ell/items/6eb48e934a147898d823>

文字の分割（split）
<https://note.nkmk.me/python-split-rsplit-splitlines-re/>

文字列をリスト関数に入れると１文字ずつ分割
<https://niwakomablog.com/python-list-string/>

先頭が小文字と大文字で処理を変える

```python
text = str(input('input text>>'))

if text[0].islower():
    print("".join([t.upper() for t in text]))
else:
    print(text * 2)
```

リストの文字列結合
    >"".join(c_list)

文字列が大文字：str.isupper()
文字列が小文字：str.islower()

整数値だけ取り出す処理

```python
l = ['1', 2, '3', 4.0, '5', 6, '7', 8.0, '9', 10]
v_list = [v for v in l if isinstance(v, int)]
print(v_list)
```

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
    <https://www.lifewithpython.com/2017/11/python-line-break-handling.html>

フォーマット　０埋め
    <https://note.nkmk.me/python-zero-padding/>

```python
    # 四則演算
    a = 7
    b = 3
    print(a + b)
    print(a - b)
    print(a * b)
    print(f'{a // b}...{a % b}')
    print(a / b)

    # 二乗
    n: int = int(input('input num plz>>'))
    print(n**2)

    # fizzbuzz
    print(['fizzbuzz' if i % 15 == 0
    else 'fizz' if i % 3 == 0
    else 'buzz' if i % 5 == 0 else i for i in range(1, 31, 1)])


Pythonで文字列リテラル'...', "..."の前にrまたはRをつけると、エスケープシーケンスを展開せずそのままの値が文字列となる。
Windowsのパスや正規表現のパターンなど、バックスラッシュを多用する文字列を扱うときに便利。
https://note.nkmk.me/python-raw-string-escape/

```

## 第2章 コレクション+zip関数+emurate関数

```python

# リスト
s_list = [1, 2, 3] 

# タプル（変更不可）
t_taple = (1, 2, 3)
t_taple = tuple(s_list)
t_tuple = (1, [2, 3], '4', (5, 6, 7), '8', (9, 10))
for n, t in enumerate(t_tuple):
    print(f'{n+1}行目 type={type(t)} value={t}')

# t = (1)はint型
t = (1,)
print(type(t))


# セット（辞書）
d1 = {'A': 111, 'B': 222, 'C': 333}

# メソッドを使った追加
d1.update({'d': 444})

# メソッドを使わない追加
d1['E'] = 555
print(f'1つ目の作成した辞書：{d1}')

# セット連結
d1 = {'A': 111, 'B': 222, 'C': 333}
d2 = {'D': 444, 'E': 555}
d3 = {'F': 666}
new_d = {}
for d in [d1, d2, d3]:
    new_d.update(d) 
print(f'連結した辞書：{new_d}')

# セット連結（開封）
d1 = {'A': 111, 'B': 222, 'C': 333}
d2 = {'D': 444, 'E': 555}
d3 = {'F': 666}
new_d = {**d1, **d2, **d3}   
print(f'連結した辞書：{new_d}')

# keyがあるか
d = {'A': 111, 'B': 222, 'C': 333}
result = 'A' in d.keys()
print(f'Aは辞書に含まれているか：{result}')

# リストからkeyセット
keys = ['A', 'B', 'C']
values = [111, 222, 333]
new_d = {}
[new_d.update({k: v}) for k, v in zip(keys, values)]
print(new_d)

# 重複除外
l1 = [1, 2, 3, 4, 5]
l2 = [7, 6, 5, 4]
r = set(l1) | set(l2)
print(r)

# l1のみの要素
r = set(l1) - set(l2)
print(f'l1のみに格納されている要素 : {r}')

# l1かl2のどちらかの要素
r = set(l1) ^ set(l2)
print(f'どちらか一方に格納されている要素 : {r}')

# 集合が含まれているか
print(f'l1はl2に含まれている : {set(l1).issubset(set(l2))}')
print(f'l2はl1に含まれている : {set(l2).issubset(set(l1))}')

# keyで並び替え
d = {'B': 222, 'A': 111, 'D': 444, 'C': 333}
sort_by_key = dict(sorted(d.items()))
print(f'keyでソートした辞書：{sort_by_key}')

# valueで並び替え
d = {'B': 222, 'A': 111, 'D': 333, 'C': 444}
sort_by_value = dict(sorted(d.items(), key=lambda x: x[1]))
print(f'Valueでソートした辞書：{sort_by_value}')

#（詳細）lambda関数 x: x[1]と同義
def get_values(x):
    return x[1]
d = {'B': 222, 'A': 111, 'D': 333, 'C': 444}
sort_by_value = dict(sorted(d.items(), key=get_values))
print(f'Valueでソートした辞書：{sort_by_value}')

# valueが奇数の要素のみ削除
d = {'B': 222, 'A': 111, 'D': 444, 'C': 333}
new_d = d.copy()
for k, v in d.items():
    if v % 2 != 0:
        del new_d[k]
        
print(f'奇数を削除した辞書 : {new_d}')


# スライド　sequence[start:stop:step]
s = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
even = s[::2]  # 偶数　0からスタートして2ステップ毎
odd = s[1::2]  # 奇数　1からスタートして2ステップ毎
print(even)
print(odd)

# リスト型のメソッドsort(): 元のリストをソート
# 組み込み関数sorted(): ソートした新たなリストを生成
```

<https://note.nkmk.me/python-key-sort-sorted-max-min/>

セットを使って重複を削除

```python
from random import randint
l = [randint(1, 9) for i in range(100)]
print(list(set(l)))
```

### 2.5　コレクションの応用

注意： ()と[]と{}を間違えないように定義しよう

set（集合型）とlistは違う
<https://aiacademy.jp/media/?p=2941>

リスト　シャッフル
<https://note.nkmk.me/python-random-shuffle/>

```python
last_names = ['aaa', 'bbb', 'ccc']

# 拡張for文
[print(name) for name in last_names if (isinstance(name, str))]

# enumerate文（自動でインデックス作ってくれる）
[print(f'i={i}, name={name}') for i, name in enumerate(last_names) if (isinstance(name, str))]

# 3の倍数にインデックスをつける
num_list = [i for i in range(1, 31, 1) if i % 3 == 0]
[print(f'count={i+1}, num={num}') for i, num in enumerate(num_list)]

# 3の倍数のインデックスと要素を削除する
l = [1, 3, 2, 3, 4, 6, 5, 8, 7]
c_list = [n for i, n in enumerate(l) if not(i % 3 == 0 or n % 3 == 0)]
print(c_list)
```

# zip関数（複数リストを）
first_list = [1, 2, 3]
second_list = [10, 20, 30, 40, 50]
third_list =[100, 200, 300, 400]
[print(first + second + third) for first, second, third in zip(first_list, second_list, third_list)]
```

入力した英単語の中央に、@を差し込んで出力するプログラム

```python
def insert_kigou(i_text):
    num_center = int(len(i_text) / 2)
    i_list = list(i_text)
    if len(i_text) % 2 != 0:
        i_list.pop(num_center)
    i_list.insert(num_center, '@')
    return i_list

a_text = 'Python'
b_text = 'PHP'

print("".join(insert_kigou(a_text)))
```

入力された文字の出現頻度をカウント

```python
text = str(input('input text>>'))
d = {}
for s in text:
    if s in d.keys():
        d[s] += 1
    else:
        d[s] = 1
print(d)

```

ランダム数列を作成して最大値と最小値を表示（max,min関数使わず）

```python
from random import randint

l = [randint(1, 99) for i in range(10)]

max_num = 0
min_num = 10000
for n in l:
    if n > max_num:
        max_num = n
    if n < min_num:
        min_num = n

print(l)
print(f'max_num={max_num} min_num={min_num}')
```

### 第3章 条件分岐

テキストaとb内で同じ文字があれば抽出する

```python
a_list = 'sample'
b_list = 'temple'

c_list = [a_list[i] for i in range(len(a_list)) if a_list[i] == b_list[i]]
print("".join(c_list))
```

```python
vowels = ['a','i','u','e','o']
word = input('文字列を入力してください >')

new_word = ''
for s in word:
    if s in vowels:
        continue
    new_word += s
    
print(f'作成した文字列：{new_word}')

```

文字列”Python”が含まれている要素を削除する
```python
import re

l = ['Python1', 'Java1', 1, 'Python2', 'Java2', 2]
pattern = 'Python'
repattern = re.compile(pattern)

r_list = [text for text in l if not repattern.search(str(text))]
print(r_list)
```

### 第4章 繰り返し

len()を使わずに文字数をカウントする

```python
def count_char(txt):
    text_list = list(txt)
    count = 0
    for _ in text_list:
        count += 1
    return count


text = str(input('input text>>'))
print(str(count_char(text)))
```

### 第5章 関数

可変長引数
<https://note.nkmk.me/python-args-kwargs-usage/>

```python
def my_sum(*args):
    return sum(args)

# 10
print(my_sum(1, 2, 3, 4))

# 36
print(my_sum(1, 2, 3, 4, 5, 6, 7, 8))
```

boolリストでFalseの数をカウント

```python
l = [False, True, False, False, True]
c = len(l) - sum(l)
print(f'Falseの数：{c}')
```

引数にリスト、タプル、辞書を展開して渡す
<https://note.nkmk.me/python-argument-expand/>

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

総まとめの演習問題

1. ジャンケン
   
2. json
>JSON文字列を辞書に変換: json.loads()
順番を保持: 引数object_pairs_hook
バイト列を変換
JSONファイルを辞書として読み込み: json.load()
読み込んだ辞書の値の取得・変更・削除・追加
辞書をJSON文字列として整形して出力: json.dumps()
区切り文字を指定: 引数separators
インデントを指定: 引数indent
キーでソート: 引数sort_keys
Unicodeエスケープ指定: 引数ensure_ascii
辞書をJSONファイルとして保存: json.dump()
JSONファイルの新規作成・更新（修正・追記）
JSONファイル・文字列を扱う上での注意点
Unicodeエスケープ
引用符
<https://note.nkmk.me/python-json-load-dump/>
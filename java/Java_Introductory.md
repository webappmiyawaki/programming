# スッキリわかるJava入門 第3版

### 第0章　Javaをはじめよう

## ■第Ⅰ部　ようこそJavaの世界へ

### 第1章　プログラムの書き方

#### エントリーポイント

```java
import main
public class Sample{
    public static void main(String[] args){
        System.out.println("hello world.")
    }
}
```

### 第2章　式と演算子

データ型

```Java
byte b = 2;
short s = 127;
int i = 10;
long num = 20000L;
float f = 1.4F;
double d = 1.3;
boolean is = true;
char c = 'c'; c =98;
String s = "String";

#定数
final int FINAL_NUM = 100;
```

++ インクリメント
-- デクリメント

型変換
小さな型から大きな型は可能(int->long)

```Java
# 文字を数字に変換
Integer.parseInt('10');

# 乱数
int r = new java.util.Random().nextInt();

# 入力
String s = new java.util.Scanner(System.in).nexlLine();
```

### 第3章　条件分岐と繰り返し

文字列の比較は`equslas`

```Java
# if文
if(a==b){
    a=1
}else{
 b=2
}

while(true){

}

```

### 第4章　配列

配列、要素、添え字

```Java
int[] scores;
scores = new int[5];
```

### 第5章　メソッド

```Java
public class Main{
    public static void main(String[] args){
        Sample sample = new Sample();
        int i = sample.methodName(1,2);
        System.out.println(i);
    }
}

class Sample{
    # 公開範囲　戻り値の型　メソッド名（シグネチャ）
    public int methodName(int1, int2){
        int i = int1 + int2;
        return i;
    }
}
```

オーバーロード（引数の型が違うメソッド）
オーバーライド（メソッドの上書き）

値渡し（プライマリー）
参照渡し（配列等）

### 第6章　複数クラスを用いた開発

パッケージ管理が必要
完全修飾クラス名（full qualified class name）

## ■第Ⅱ部　すっきり納得オブジェクト指向

### 第7章　オブジェクト指向をはじめよう

### 第8章　インスタンスとクラス

### 第9章　さまざまなクラス機構

コンストラクタ

- メソッド名とクラス名が同じ
- メソッド宣言に戻り値が記述されていない
- 記述されていない場合も引数なし処理なしで自動で追加される

```Java
class con{
    con(){
        # コンストラクタ        
    }
}
```

thisとsuperの使い分け

```Java
#自クラスのメソッド
this.method;

#親クラスのメソッド
super.method;

```

### 第10章　継承

extends

### 第11章　高度な継承

抽象クラス

```Java
public abstract class Character{
    # 処理内容未設定
    public abstract void attack(Mattango m);
}
```

インターフェース

- 全てのメソッドが抽象メソッド
- フィールドを持たない
- implementsで実装


```Java
interface Inter{
    interMethod();
}

public class Outer implements Inter{

}
```

### 第12章　多態性

ポリモーフィズム（polymorphism）

### 第13章　カプセル化

- private
- setterメソッド
- getterメソッド

アクセス修飾詞
public >> package private >> protected >> private

## ■第Ⅲ部　もっと便利にAPI活用術

### 第14章　Javaを支える標準クラス

親クラスを指定しない場合`java.lang.Object`を承継しているとみなされる

`toString()`と`equals()`のオーバーライド

### 第15章　文字列と日付の扱い

Stringクラスのメソッド

- contains（一部に含むか）
- startsWith（文字列sで始まるか）
- endsWith（文字列sで終わるか）
- indexOf
- lastIndexOf
- charAt（１文字切り抜き）
- substring（指定した場所から任意の文字数切り出す）
- toLowerCase（小文字にする）
- toUpperCase（大文字にする）
- trim（前後の空白を除去）
- replace（文字列を置き換える）
- matches（正規表現マッチ）
- split（文字列分割）

StringBuilderが便利

### 第16章　コレクション

list,set,map

listは基本データ型を使えないが、
基本 <--相互変換-->ラッパー型
できる。（オートボクシング、オートアンボクシング）

### 第17章　例外

- Error系（回復見込みなし。catchする必要なし）
- Exception系（catchすべき）
- RuntimeException系（catchしなくてよい）


```Java
try{
    処理
}catch(Exception e){
    例外時に行う処理
}finally{
    絶対実行される処理
}
```


### 第18章　まだまだ広がるJavaの世界

付録A　ローカル開発環境のセットアップと利用
付録B　エラー解決・虎の巻
付録C　クイックリファレンス
# vscodeでflask

1. venvによる仮想環境の作成

自分の使いたいフォルダで仮想環境作成
-scriptフォルダ等作成される

すでにフォルダがある場合
venv（仮想化用に使われるフォルダ）をvenvコマンドで作る

```console
python -m venv venv 
```

今から作るプロジェクト

```console

```

requirements.txt（フォルダに注意）
中身（2022/05/05）

```Text
click==8.1.3
colorama==0.4.4
Flask==2.1.2
Flask-Login==0.6.1
Flask-SQLAlchemy==2.5.1
greenlet==1.1.2
importlib-metadata==4.11.3
itsdangerous==2.1.2
Jinja2==3.1.2
MarkupSafe==2.1.1
pytz==2022.1
SQLAlchemy==1.4.36
Werkzeug==2.1.2
zipp==3.8.0
```

```console
# 設定書き出し時
pip freeze > requirements.txt

# 設定読み込み時
pip install -r requirements.txt
```


仮想環境を有効にするには、ステップがもう1つ必要です。ステータスバー左端のPythonバージョンをクリックするか、コマンドパレットから［Python: インタープリターを選択］コマンドを実行して、作成した仮想環境にあるPythonインタプリターを選択します

vscodeの設定＞＞
python.terminal.activateEnvInCurrentTerminal項目にチェックを入れる

## VSCodeでActivate時に「このシステムではスクリプトの実行が無効になっているため、ファイルｘｘを読み込むことができません」とエラーが表示される時の対処法

<https://atmarkit.itmedia.co.jp/ait/articles/2106/25/news028.html>

VSCodeの setting.json を開き、以下のように記述しちゃいましょう。

```console
"terminal.integrated.env.windows": {
        "PSExecutionPolicyPreference": "RemoteSigned"
    }
```

（flask-sqlalchemyが必要かも）
>pip install flask-sqlalchemy

対話モード実行
>python

app.py内で定義したflask_sqlalchemyのインスタンスであるdb変数をインポート
>from app import db

app.py内のBlogArticleクラスに対応したDBのテーブルが作成
>db.create_all()

2. Flaskのインストール基本構成

```Text
(app名)
static（css,javascrip用フォルダ）
    - css
    - javascript
templates
    - base.html
    - create.html
    - index.html
    - login.html
    - signup.html
    - update.html 
venv(仮想化用のexeが入ってるフォルダ)
app.py（実行ファイル）
blog.db（データ登録用データベース）
requirements.txt（設定読み込み用ファイル）
```

3. コードの記述


4. デバッグ実行
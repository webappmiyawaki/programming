# vscodeでflask

1. venvによる仮想環境の作成

自分の使いたいフォルダで仮想環境作成
-scriptフォルダ等作成される

すでにフォルダがある場合

```console
python -m venv venv 
```

今から作るプロジェクト

```console

```

requirements.txt（フォルダに注意）

```console
pip freeze > requirements.txt
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
>python3

app.py内で定義したflask_sqlalchemyのインスタンスであるdb変数をインポート
>from app import db

app.py内のBlogArticleクラスに対応したDBのテーブルが作成
>db.create_all()



2. Flaskのインストール


3. コードの記述


4. デバッグ実行
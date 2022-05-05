# flask

flaskでwebアプリ
    <https://engineer-lifestyle-blog.com/code/python/flask-usage-tutorial-for-beginner/>

flask基本
<https://qiita.com/bauer/items/70abcb68d3b00d0d1794>

■ 関連ドキュメント
Flask
<https://flask.palletsprojects.com/en/>...

Flask Login
<https://flask-login.readthedocs.io/en>...

Bootstrap
<https://getbootstrap.jp/>

sqlalchemy
<https://flask-sqlalchemy.palletsprojects.com/en/2.x/>

>flask_sqlalchemyはsqlalchemyとは別のモジュールです。
pip install flask-sqlalchemy
でインストールしてください。

Flask-SQLAlchemyの使い方

requirements.txt
<https://qiita.com/sakusaku12/items/21083c73c8afa4f6c78d>

YouTube
<https://www.youtube.com/watch?v=VtJ-fGm4gNg>

### ルーティング

    アドレスを可変にする（型宣言可）

    ```python
    @app.route("/japan/<city>")
    def japan(city):
        return f'hello, {city} in Japan!'
    ```

### テンプレートを用いる（Jinja）

    #render_templateが必要

    ```python
    from flask import Flask
    from flask import render_template

    app = Flask(__name__)


    @app.route('/')
    def hello():  # put application's code here
        return render_template('hello.html')
    ``` 

    # セッション登録（html側で呼び出しは`{{html_city}}`） 

    ```python
    @app.route('/japan/<string:city>')
    def hello(city):  # put application's code here
        # シグネチャ(html_cityはセッション名、cityは変数)
        return render_template('hello.html', html_city=city)
    ```

    # html側でリストを呼び出す
    # for-endfor　が必要
    # enumrationが使えないのでloop.indexを代わりにする

    ```html
    <h1>サンプル{{ html_city }}</h1>
    <ul>
        {% for b in num_list %}
            <li>{{ b, loop.index }}</li>
        {% endfor %}
    </ul>
    ```

＜Flask vol.2 ブログアプリの作成〜DB連携〜＞
要素の共通化

    ```html
    #base.html側（共通化用パーツ）
    {% block content %}
    {% endblock %}

    #index.html側（表示ページ側）
    {% extends "base.html" %}
    {% block content %}
        ～省略～
    {% endblock %}
    ```


DB作成
新規登録
表 示
編 集
削 除

＜Flask vol.3 ログイン機能の実装とCSS適用＞
概要説明
Flask-Loginのインストール
ユーザーDBの作成
サインアップ機能
ログイン機能
ログアウト機能
動作確認
CSS, JSを適用
Bootstrapの適用
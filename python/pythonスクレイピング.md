# python

beautiful soapのインストール（pipがない場合）

```console
python -m pip install --upgrade pip
pip install requests
pip install beautifulsoup4
```

```python
from bs4 import BeautifulSoup
from urllib import request

url = 'https://www.atmarkit.co.jp/ait/subtop/di/'
response = request.urlopen(url)
soup = BeautifulSoup(response)
response.close()

topstories = soup.find('div', class_='colBoxTopstories')
print(topstories)

colboxindexes = topstories.find_all('div', class_='colBoxIndex')
print(colboxindexes[0])

title = colboxindexes[0].select('div.colBoxTitle')[0].text
description = colboxindexes[0].select('div.colBoxDescription')[0].text
print(title, ':', description)

top_articles = []
for item in colboxindexes:
    title = item.select('div.colBoxTitle')[0].text
    description = item.select('div.colBoxDescription')[0].text
    top_articles.append(f'{title}: {description}')

for articles in top_articles:
    print(articles, end='\n')
```
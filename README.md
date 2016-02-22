# 概要
Google App Engine で静的コンテンツを配信するための雛形
※まだ作成中です！！
 
## 目的

このプロジェクトは、「Google App Engine for Go」で静的コンテンツを配信するためのサンプルです。

Google App Engineの入門用として自由にご利用下さい。

## 使い方

### Google App EngineのSDKインストール

GolangのSDKをインストールします。

https://cloud.google.com/appengine/downloads

### 当プロジェクトをgit clone

```
git clone ....
```

### staticディレクトリ内に静的コンテンツを作成する。

app/staticの中身を作成して下さい。

notfound.htmlは、名前を変更せずに内容を変更して下さい。

### ローカルでの実行
```
cd [[project_home]]
goapp serve 
```

### プロダクト環境へのデプロイ

app.yamlの `application` を自分のIDに変更
```
goapp deploy
```
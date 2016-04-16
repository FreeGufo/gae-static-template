# 概要
Google App Engine で静的コンテンツを配信するための雛形
 
## 目的

このプロジェクトは、「Google App Engine for Go」で静的コンテンツを配信するためのサンプルです。

Google App Engineの入門用として自由にご利用下さい。

## 使い方

### Google App EngineのSDKインストール

GolangのSDKをインストールします。

https://cloud.google.com/appengine/downloads

※ SDKのインストールの公式にもありますが、GoのSDKを使用する際もPython2.7が必要になります。

### staticディレクトリ内に静的コンテンツを作成する。

app/staticの中身を作成して下さい。

notfound.htmlは、名前を変更せずに内容を変更して下さい。

設定項目の詳細は、こちらを参照して下さい。

http://go-talks.appspot.com/github.com/FreeGufo/gae-static-template/go-talks/index.article

当リポジトリの「go-talks」ディレクトリにも同じ内容が入っています。

### ローカルでの実行
```
cd [[project_home]] ※ app.yamlのあるディレクトリ
goapp serve 
```

### プロダクト環境へのデプロイ

Cloud Consoleでプロジェクトを作成
https://console.cloud.google.com/


app.yamlの `application` を作成した自分のIDに変更
※ Project名とProjectIDは、別に付けられるので注意。IDを使用して下さい。

詳しくは、公式サイトを参照して下さい。
https://cloud.google.com/appengine/docs/go/

ここまで、Product側とローカルの準備ができたら以下を実行

```
cd [[project_home]] ※ app.yamlのあるディレクトリ
goapp deploy
```

### アクセス方法

デプロイが成功していれば、インターネット経由でアクセスが出来るようになっているはずです。
URLは、以下の様に、 `appspot.com` のサブドメインとなります。

http://[[applicationid]].appspot.com/


# gae-test

GAE で Go アプリケーションを起動するテストアプリです。

## 起動し動作を確認

アプリは、GET リクエストを受信し、conf.json の内容を JSON で返却するだけのものです。<br/>
基本の API サーバーを GAE で動作させることが当アプリの目的です。

### アプリの起動

#### go run main.go

GAE のデプロイするための設定を確認してください。<br/>
Google Cloud Platform SDK のインストール済みを想定しています。<br/>
もしインストールがまだの場合は<a href="https://qiita.com/G-awa/items/e6904b040caa0096fba0">この記事</a>を参考にインストールしてください。<br/>
また、<a href="https://cloud.google.com/sdk/gcloud?hl=ja0">SDK の公式ドキュメント</a>も参考にしてください。

### GAE に関する注意点

GAE には Standard Environment と Flexible Environment という２つの実行環境があります。<br/>
詳しくは<a href="https://cloud.google.com/appengine/docs/the-appengine-environments">この記事</a>を参考にしてください。<br>
チュートリアルは <strong>Flexible Environment</strong> を採用していますが、GAE の無料枠の対象となるのは <strong>Standard Environment</strong> のみのため、注意が必要です。<br/>
さらに、デフォルトではインスタンスが２台立ち上がるようになっているため、知らず知らずのうちに課金額が大きくなってしまう、ということがあります。

### SDK の確認

#### $ gcloud config list

### プロジェクトの変更

#### $ gcloud config set project PROJECT_NAME

### アプリのデプロイ

app.yaml を<a href="">この記事</a>を参考に作成しコマンドを実行してください。

#### $ gcloud app deploy

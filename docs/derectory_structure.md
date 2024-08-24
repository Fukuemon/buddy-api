# ディレクトリ構成

```
.
├── README.md
├── changelog.config.js // git cz用ファイル
├── docker // DockerFile
├── docker-compose.yml
├── script // DB構築, 初期データ投入スクリプトなど
├── docs
│   └── derectory_structure.md
└── internal // 下記記載
    ├── cmd
    │   └── api-server
    │       └── main.go
    ├── config
    ├── domain
    ├── go.mod
    ├── infrastructure
    ├── presentation
    ├── server
    │   └── route
    └── usecase
```

# internal

アプリケーションに関するファイル

## cmd

アプリケーションのエントリーポイントやマイグレーション用のファイル

## config

環境変数など設定値を定義するファイル

## server

HTTP サーバーの設定や、ルーティングの設定を行うファイル群

### server/route

API のエントリーポイントを定義するファイル

## Entity Layer

### domain

- Domain Model
- Domain Service

集約単位でディレクトリを切って管理する。
それぞれ Model・Service を定義する

#### Domain Model

ドメインオブジェクトを定義し、ビジネスルールを実装する
例：

- 名前の長さは〇〇字まで
- 数値の不正など

#### Domain Service

単一のドメインオブジェクトの責務だけでは処理できないビジネスロジックを担う
例：注文処理（product・order・cart）

- 商品の注文
- 履歴の保存

この処理は、注文の集約で担うべきだが、処理の中で、
商品・注文・カートのドメインオブジェクトが必要になる。
この場合、注文のドメインオブジェクトではビジネスルールを表現できないため、
ドメインサービスとして他のドメインオブジェクトを扱って実装する。

## InterFaceAdapter Layer

### infrastructure

- QueryService
- Repository

永続化層の処理を担当する
mysql・Cognito の設定ファイルや repository の実装

### presentation

- Handler

request・response の構造体を定義し、各ドメインのハンドラーを実装する

## UseCase Layer

### usecase

Entity Layer のオブジェクトや関数を用いて、ユースケースの処理を実行する
domain と同じようにディレクトリを構成する
各ドメインでは、ユースケースごとにファイルを分けて実装する

# API 仕様書（Swagger）

[swaggo](https://github.com/swaggo/swag)と[gin-swagger](https://github.com/swaggo/gin-swagger)を利用して、API 仕様書の WebPreview と仕様書ファイル（.yml/.json）の自動生成を行う

### Web での閲覧方法

（アプリを起動しているものとする）

1. [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)にアクセス

### 定義ファイル

`internal/docs ｀の中に格納されている

```
docs
├── docs.go // swaggo設定ファイル
├── swagger.json // json形式の仕様書
└── swagger.yaml // yaml形式の仕様書
```

アプリのコンテナ起動時に自動で最新のものに更新される。
手動で更新する場合は、下記のコマンドを実行する

```
docker-compose run --rm update_api_docs
```

### Postman のコレクション自動生成・更新

`.env`に下記の環境変数を登録

```
POSTMAN_API_KEY=<<api_key>>
POSTMAN_COLLECTION_ID=<<collection_id>>
```

下記のコマンドを実行すると、対象のコレクションが自動で更新される

```
docker compose run --rm update_postman
```

# 🐱フロントエンド開発を頑張りたい人に向けた練習用バックエンドAPI🐱

## 起動前に
- dockerとgoをお使いのPCにインストールしてください
- env.exampleの名前を.envに変更してください
- 各エンドポイントの詳細はSwaggerに記載しています

## PORT

### API

```
localhost:18080
```

### Swagger

```
localhost:18081
```

### MySQL

```
localhost:13306
```


## コマンド

### docker 起動

```sh
docker-compose up -d
```

### docker 停止

```sh
docker-compose down
```

### API 起動

```sh
go run main.go
```

---

## 以下 API 開発用コンソール

### UseCase 作成

```sh
node console.js u ${dirName} ${useCaseName}
```

### Repository 作成

```sh
node console.js r ${repositoryName}
```

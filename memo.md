### http server作成 p127

- main.go
- ポート開ける
- hello world

### リファクタリングとテストコード

テスト容易性

- run関数に分離する

- エラーパッケージ


```
go get -u golang.org/x/sync
```

- eg.Go ゴルーチンの並列管理を簡潔にできるらしい

- ⭐️テストコードはスキップする


### ポート変更可能に

net.Listenerを利用

コマンドで実行
```
go run main.go 8080
```

### 開発環境を整える p.140

dockerを利用

.dockerignore

- ⭐️マルチビルドのポイント！！！

docker-composeで devを見たい場合は、targetで指定できる

```
version: "3"
services:
  app:
    build:
      context: .
      target: dev  # ← ここで dev ステージを指定
    volumes:
      - .:/app
    ports:
      - "8080:8080"
    command: ["air"]
```

- dockerコマンド

```
// ビルド
docker-compose build --no-cache

// 起動 target devなのでdev指定

docker-compose up

localhost:1800/hello
```


### makefile p.145

コピーした

### github actions p147 スキップ

### httpサーバー疎結合にする

#### 環境変数から設定をロードする p152

port とかをハードコーディングしてる点が、依存度が高いのか

- 環境変数
  - 記事 何パターンかある
    - https://zenn.dev/kurusugawa/articles/golang-env-lib
  - 全体像
    - ライブラリ入れる
    - 環境変数を参照する
    - docker-composeで環境変数を渡す
      - envファイルを作成して参照でもできた気がする

```
go get github.com/caarlos0/env/v10
```

- configを作成した
  - New関数で構造体を返却して利用
    ```
      cfg := &Config{} // これはデフォルトがあるので何も引数に渡していない
    ```

### シグナルをハンドリングする


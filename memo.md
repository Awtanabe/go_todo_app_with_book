### コード

https://github.com/budougumi0617/go_todo_app/tree/main

### イメージ

- ディレクトリ
  - entity
    - ドメインを
  - auth
    - 認証
    - ミドルウェアで
  - config
    - 環境変数
  - handler
    - アクション
  - service
    - usecase
    - interface
  - store
    - これは miidasで言うところの xx_repo.go
    - repogitory.go
      - 実態
    - user.go
      - レシーバーを持った関数
- 暗号化
  - crpt
- 認証
  - jwt
    - ログインする
    - tokenを発行する。http onlyでクッキーに保持させる
    - アクセス制限
      - ミドルウェアでチェック
        - jwtとredisのデータと照合して判定
      - contextに、ユーザーの情報を引き回す
- mux理解
  - main, server, mux理解する
  - http/netパッケージでサーバー
    - https://qiita.com/BitterBamboo/items/6119f7986a04c5a0ac57
  - chiはルーティングのミドルウェアらしい
    - https://github.com/go-chi/chi
- テストコード

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


### Server構造体を定義する p159

server.go

### ルーティング定義を分離したNewMux p162

mux.go

### run関数のリファクタリング p164

### エンドポイントを追加 p165


#### entuty.Task型の定義と永続化方法の仮実装(dbを使わない)


- defined type
https://zenn.dev/nobishii/articles/defined_types

⭐️使うシーン
```
	var id int = 1

   // ok
	_ = Task{ID: TaskID(id)}
  // ng 別の型で定義されているから
	_ = Task{ID: id}

  // ok  リテラルはOK！！ 型推論
	_ = Task{ID: 1}
```


### ⭐️データ型の理解


- type Task []*Task
  - Task は *Task のスライス

```
package main

import (
	"encoding/json"
	"fmt"
)

// Task はタスクの構造体
type Task struct {
	Name     string  `json:"name"`
	SubTasks []*Task `json:"sub_tasks"`
}

func main() {
	// タスクの作成
	rootTask := &Task{Name: "Root Task"}

	// サブタスクを作成
	subTask1 := &Task{Name: "Sub Task 1"}
	subTask2 := &Task{Name: "Sub Task 2"}

	// サブタスクの中にさらにサブタスクを追加
	subTask1_1 := &Task{Name: "Sub Task 1-1"}
	subTask1_2 := &Task{Name: "Sub Task 1-2"}
	subTask2_1 := &Task{Name: "Sub Task 2-1"}

	subTask1.SubTasks = []*Task{subTask1_1, subTask1_2}
	subTask2.SubTasks = []*Task{subTask2_1}

	// ルートタスクにサブタスクを追加
	rootTask.SubTasks = []*Task{subTask1, subTask2}

	// JSON 形式で出力
	jsonData, err := json.MarshalIndent(rootTask, "", "  ")
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return
	}
	fmt.Println(string(jsonData))
}

```

```
package main

import (
	"encoding/json"
	"fmt"
)

// Task はタスクの構造体
type Task struct {
	Name     string  `json:"name"`
	SubTasks []*Task `json:"sub_tasks"`
}

func main() {
	// ルートタスク作成
	rootTask := &Task{Name: "Root Task"}

	// サブタスク作成
	subTask1 := &Task{Name: "Sub Task 1"}
	subTask2 := &Task{Name: "Sub Task 2"}

	// `append` を使って子タスクを追加
	rootTask.SubTasks = append(rootTask.SubTasks, subTask1, subTask2)

	// さらに `subTask1` にサブタスクを追加
	subTask1_1 := &Task{Name: "Sub Task 1-1"}
	subTask1_2 := &Task{Name: "Sub Task 1-2"}
	subTask1.SubTasks = append(subTask1.SubTasks, subTask1_1, subTask1_2)

	// `subTask2` にもサブタスクを追加
	subTask2_1 := &Task{Name: "Sub Task 2-1"}
	subTask2.SubTasks = append(subTask2.SubTasks, subTask2_1)

	// JSON 形式で出力
	jsonData, err := json.MarshalIndent(rootTask, "", "  ")
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return
	}
	fmt.Println(string(jsonData))
}

```

### ヘルパーの実装⭐️ p169

- unmarshal
  - jsonを扱う感じのやつだった気がする


### タスクを登録する ⭐️172p(バリデーション)


- バリデーション
  1. ライブラリを入れる
  2. AddTask構造体を作成、バリデーターをDIで保持(バリデーションの実行)
  2. paramsデータの構造体を定義する(requredなどのデータのチェック)
  3. jsonDecoderでチェク
    ```
      // r.Bodyでrequest bodyを取得。decodeで解析かな？
      if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
        return errors.New(err.Error())
      }
      // 処理をする
      return 最終的なレスポンス
    ```

- データを入れる
```
	t := store.Tasks
	t.Tasks[0] = &entity.Task{ID: 1, Title: "テスト", Status: entity.TaskStatusDosing, Created: time.Now()}
	lt := &handler.ListTask{Store: t}

```

### http ハンドラーをルーティングに設定 p179

```
go get -u github.com/go-chi/chi/v5

chi.NewRouterで利用
```

### mysql 環境構築 p184

#### 必要なもの

- パッケージ

```
database/sql

// gorm

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql



go get -u github.com/go-sql-driver/mysql

go install github.com/k0kubun/sqldef/cmd/mysqldef@latest
```

- mysql.cnf

```
[mysql]
default_character_set=utf8mb4

```

- mysqld.cnf

```
[mysqld]
default-authentication-plugin=mysql_native_password
character_set_server=utf8mb4
sql_mode=TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY

```

- compose

```
TODO_ENV: dev
PORT: 8080

TODO_DB_HOST: todo-db
TODO_DB_PORT: 3306
TODO_DB_USER: todo
TODO_DB_PASSWORD: todo
TODO_DB_DATABASE: todo

volumes:
  - .:/app


todo-db:
  image: mysql:8.0.29
  platform: linux/amd64
  container_name: todo-db
  environment:
    MYSQL_ALLOW_EMPTY_PASSWORD: "yes"
    MYSQL_USER: todo
    MYSQL_PASSWORD: todo
    MYSQL_DATABASE: todo
  volumes:
    - todo-db-data:/var/lib/mysql
    - $PWD/_tools/mysql/conf.d:/etc/mysql/conf.d:cached
  ports:
    - "33306:3306"

volumes:
  todo-db-data:

```

- sql

```
CREATE TABLE `user`
(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`      VARCHAR(20) NOT NULL COMMENT 'ユーザー名',
    `password`  VARCHAR(80) NOT NULL COMMENT 'パスワードハッシュ',
    `role`      VARCHAR(80) NOT NULL COMMENT 'ロール',
    `created`   DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified`  DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `task`
(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `title`     VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `status`    VARCHAR(20) NOT NULL COMMENT 'タスクの状態',
    `created`   DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified`  DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';

```

### 動作確認


```
 curl -i -X GET http://localhost:18000/tasks

 curl -i -X POST -H "Content-Type: application/json" -d '{"title":"タスク4"}' http://localhost:18000/tasks

  curl -i -X POST -H "Content-Type: application/json" -d '{"title":"タスク6"}' http://localhost:18000/tasks


 curl -i -X POST -H "Content-Type: application/json" -d '{"name":"マイク", "password":"password", "role": "管理者"}' http://localhost:18000/register

  curl -i -X POST -H "Content-Type: application/json" -d '{"name":"マイク", "password":"password"}' http://localhost:18000/login

```

### ハンドラーとビジネスロジックを分離する p206

service層


### ユーザー登録 p214


### p223 Redisの準備

- 作業
  - composeで起動
	- アプリケーションコード
	  - store/kvs.goを作成
		- 環境変数

### jwtで行う署名の準備 p230

- 秘密鍵、公開鍵

openssl genrsa 4096 > secret.pem
openssl rsa -pubout < secret.pem > public.pem

- jwtを用いてアクセストークン作成

json web token
=> 署名(rs256)と暗号化に関わる関連仕様がある

https://qiita.com/asagohan2301/items/cef8bcb969fef9064a5c#2-jwt%E8%AA%8D%E8%A8%BC%E3%81%A8%E3%81%AF

- jwtとは
トークンベース認証
=> ユーザーはログイン情報を保持しない
=> 毎回のリクエストに一緒に送る
=> 改ざんに検知できる

- 作り方
=> ヘッダ.ペイロード.署名

=> 前半(ヘッダ.ペイロード)
ヘッダ: アルゴリズムとトークンタイプ
ペイロード: sub(ユーザーを一意に特定するための識別子), exp: 有効期限
※パスワードなどは保持しないs
=> 後半(署名)
指定したアルゴリズムで署名する

- jwtの検証

- jwt認証の流れ

1. ユーザー: 情報をpost
2. API: jwtの署名 && OKからjwtを発行しレスポンス
3. jwtで検証する
※ユーザーはjwtをリクエストヘッダーに自動で含める => 色々あるけど
=> HttpOnly Cookieが楽だろな (ブラウザーのcookieに入れる)

- jwt ログインイメージ
  - ログインする
  - jwttokenをクッキーで有効期限がなくなるまで使う
  - 他のエンドポイントでtoken利用
    - ミドルウェアで検証

- 必要なパッケージ

```
go get github.com/lestrrat-go/jws/v2
go get github.com/google/uuid
```

### context.Context でjwtを引き回す p241

jwt.goに実装


### repogitory

// repogitoryの実態
store/repogitory.go


// store.go
// user.go でクエリをかく

// handler/service.go
interfaceを書いてる => ここじゃなさそう

### エンドポイントの保護

- taskテーブル

```
CREATE TABLE `task`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `user_id`    BIGINT UNSIGNED NOT NULL COMMENT 'タスクを作成したユーザーの識別子',
    `title`      VARCHAR(128) NOT NULL COMMENT 'タスクのタイトル',
    `status`     VARCHAR(20) NOT NULL COMMENT 'タスクの状態',
    `created`    DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified`   DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
        ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';

```

### 動作確認 adminはスキップ


```
// user作成
curl -X POST localhost:18000/register -d '{"name": "normal_user", "password":"password", "role":"user"}'

// ログイン
curl -X POST localhost:18000/login -d '{"name": "normal_user", "password":"password"}'

// アクセストークンをローカルに
export TODO_TOKEN=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDE4MDEwMTksImlhdCI6MTc0MTc5OTIxOSwiaXNzIjoiZ2l0aHViLmNvbS9idWRvdWd1bWkwNjE3L2dvX3RvZG9fYXBwIiwianRpIjoiYzY4ZWJiYjItNDUyOS00ZTU0LThkNmItMzAxNzIwNWE1NzQ1Iiwicm9sZSI6InVzZXIiLCJzdWIiOiJhY2Nlc3NfdG9rZW4iLCJ1c2VyX25hbWUiOiJub3JtYWxfdXNlciJ9.fgAHJjNk1xI8jXW0lhar9Fk21oAQDtVhhIVSM2ffWrvCGwWLdfdTzXHaoEvZtogg8_A765u9gONDOwH327Ko74sJvb3l_Uh1KB8JJyFrI3M4ZfPpl-THMevPWHjBmlD0DnBZqmHUSzsFSZ7-3S5uS--CAkleFXK3q27_sBAhA-O-fKY-6syGTSE6AcwgXUEyXBLgG2B4T05KjP2Px3uoZG7VoiuHBDPJ-Kfe25VKHDDvqQGWLHPwEwl00T5NpqaVwzZNqbmpBST_I50vcDeJVo1IOXisoPfVpoEpEv852iEXlxbLBNkabReCXKKh9L6tLLdAHoDfqeDn9Z8gnHF3-9WvH69QRScTd690ipsjDWqkkXQKHOZFLBpiihJGhlfho8YNJ-ZhM6MmM4a9SWf17jmw0iznHy8Y54TrRjLhuhD4qm8i_GSrU4zwuPdGm-Lx41gPLgWZxDRChfqlrUjLP_smQ-8X3A1RlKAGrXUpGmaC4b4MhjuDYEkFH9DmyedcG6QxoblUeQXBO2taZSKClVFqiZGPAOkYSqgqZ1gCr2ShDvwMZUttLejsmp__M9amqaMrNb-pWvJiivh6XZsAnIGIx6arivSbtkp0ZBh6ua1LuUAoT82JIbfO7fkJFhJiKZO9G_y7bn4HjBGJCkYizxerQLotRbP4-P7N44MKMJU

// 登録

curl -X POST -H "Authorization: Bearer $TODO_TOKEN" localhost:18000/tasks -d '{"title": "タスク1"}'
curl -X POST -H "Authorization: Bearer $TODO_TOKEN" localhost:18000/tasks -d '{"title": "タスク2"}'


// 作成
curl -X GET -H "Authorization: Bearer $TODO_TOKEN" localhost:18000/tasks

curl -X GET localhost:18000/tasks
```
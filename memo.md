### コード

https://github.com/budougumi0617/go_todo_app/tree/main

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

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
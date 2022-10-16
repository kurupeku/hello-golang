# ナ ◯ アツ関数

3 の倍数と 3 のつく数字だけアホになる関数を実装せよ。

1 から引数として受け取った数字 `n` までの数字を順番に文字列化して `[]string` として返すが、文字列化する数字が上記に該当する場合はアホな感じにすること。

## 前提条件

### 引数

`0 < n < 1000`

### 読み上げ関数

`func Nabeatsu(n int) []string` を実装して本課題をクリアしてください。

### 数字の文字列化

各数字の文字列化に関しては、定義済みの `type AhoNumber uint` の定義済みのメソッド `func (AhoNumber) Call()` を実装して実現してください。

e.g.)

```go
aho := AhoNumber(3)
aho.Call() // => "さぁん"
```

### 数字のアホ化

`helper.AhoString(i int)` に整数を渡すと問答無用で数字をアホ化してくれます。

全部に使うと 3 の倍数や 3 のつく数字以外もアホになるので、条件に一致するものだけ渡すようにしてください

## Tips

### 文字列変換

単純な整数の文字列化は `strconv.Itoa(i int) string` を使うと実装できます。

その他の型からの変換用関数や文字列からそれらの型への変換もこのパッケージに網羅されているので、文字列絡みの単純な変換を行う場合はまずこのパッケージから探しましょう。

### 特定の文字列が含まれるかの判定

数字のままでは特定の数字が含まれるかどうかを判定する手段がないので、文字列に変換後、文字列検索を行って判定することで実現可能です。

文字列の検索や操作用のパッケージである `strings` パッケージをインポートして活用してみましょう。

## テスト実行コマンド

1. ローカル or DevContainer 内で実行する場合
`task test -- subject/chapter06`
1. ローカルから Docker を起動して実行する場合
`task docker-build` (初回のみ)
`task docker-test -- subject/chapter06`

## CI結果確認

GitHubにpush後、以下のページで自分のブランチの結果を確認
[Testing Chapter 06](https://github.com/kurupeku/hello-golang/actions/workflows/chapter06_test.yml)

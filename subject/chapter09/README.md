# 山手線の改札機を実装せよ

切符と IC カードを利用可能な改札機を実装せよ。

距離による料金は [課題 03](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter03) と同じだが以下の機能を追加する必要がある。

- 切符と IC カードで料金が異なる
- 乗車駅と降車駅を指定して正しい料金を算出できる
- 外回りと内回りの内、距離の短い方を料金算出に適応する

## 前提条件

### 改札

`exam.go` に `func Kaisatsu(from, to string, charger Charger) (bool, error)` として定義しています。

#### 引数 `from` / `to`

- `From`: 乗車駅の名前
- `To` : 降車駅の名前

どちらの駅名も `駅` を含まない名前になります。

#### 引数 `Charger`

切符か IC カード を受け取る interface 型の引数です。

**`Ticket` を使用した際は `Used: true` に、 `Card` を使用した際は残高を減らす必要があります。**

以下のメソッドを持ちます。

- `Amount() int`
- `Use(charge int)`

内部でこの interface の実体の型を判定し、 `Ticket` と `Card` で処理を分ける必要があります。

#### 戻り値

乗車区間と乗車券に基づいて処理した結果を以下の通りに返してください。

- `bool`: 残高 or 金額不足で通れない場合は `false` を、通れる場合は `true` を返す
- `error`: 利用金額が `0` だった場合のみ `error` を、そうでなければ `nil` を返す(金額不足の場合も`nil`を返却)

`error` は 標準パッケージ `errors` の `errors.New` を使用してください。エラーメッセージに指定はありませんので何かしらそれっぽい文章を入れてください。

### 切符と IC カード

それぞれ事前に `Ticket（切符）` `Card（ICカード）` として `ticket.go` に定義済みです。

改札側には `Charger` という interface として渡すようにしているため、それらを満たすように実装を追加してください。

それぞれの細かい仕様は以下の通りです。

#### 料金表

| 下限 (m) | 上限 (m) | 切符料金 (円) | IC カード料金 (円) |
| -------: | -------: | ------------: | -----------------: |
|        - |     3999 |           140 |                136 |
|     4000 |     6999 |           160 |                157 |
|     7000 |    10999 |           170 |                168 |
|    11000 |    15999 |           200 |                198 |
|    16000 |    20999 |           270 |                264 |
|    21000 |    25999 |           350 |                341 |
|    26000 |    30999 |           420 |                418 |
|    31000 |        - |           490 |                484 |

#### Ticket

切符は購入額と使用済みか否かのみを判定できるようにしてください。

- `Price`: `>0` の `int` を購入額として記録するフィールド
- `Used` : 使用済みの場合は `true` 、未使用の場合は `false` となるフィールド

#### Card

[課題 04](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter04) と同じ仕様です。

ポイント残高があれば先に消費し、残額をチャージ残高から差し引くように実装を追加してください。

- `Balance`: `>=0` の `int` をチャージ残高として保持するフィールド
- `Point` : `>=0` の `int` を保有ポイントとして保持するフィールド

#### Charger

`Ticket` と `Card` をまとめて改札に渡す際に用いる interface です。

以下のメソッドを持ちます。

- `Amount() int`: 乗車料金として利用可能な額を `int` 型で返すメソッド
- `Use(charge int)`: 乗車料金 `charge` を `int` 型で受け取り、自身を使用した際の処理を行うメソッド

### Fare

乗車駅と降車駅を記録し、それを基に料金計算を行う構造体です。

- `From`: 乗車駅の名前
- `To` : 降車駅の名前

改札関数で受け取る引数をそのまま受け取る想定です。

内回りと外回りの距離をそれぞれ計算し、距離の短い方を基に料金を算出する必要があります。

`Ticket` と `Card` で料金が異なるため、最終的に

- `func (f *Fare) TicketCharge() int`: 切符の料金を計算するメソッド
- `func (f *Fare) CardCharge() int` : IC カードの料金を算出するメソッド

それぞれを改札関数の中で使い分けるような実装にする必要があります。

## テスト実行コマンド

1. ローカル or DevContainer 内で実行する場合

   `task test -- subject/chapter09`

1. ローカルから Docker を起動して実行する場合

   `task docker-build` (初回のみ)
   `task docker-test -- subject/chapter09`

## CI 結果確認

GitHub に push 後、[こちらのページ](https://github.com/kurupeku/hello-golang/actions/workflows/chapter09_test.yml)で自分のブランチの結果を確認

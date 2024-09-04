# hello-golang

Subjects for learning golang on the "A Tour of Go"

## Pages

- https://kurupeku.github.io/hello-golang/

## Get Started

### 実行環境の構築

以下のいずれかの方法で実行環境を構築してください。

#### ローカル環境

Go のインストールが必要です。
以下のリンクからインストール方法を確認してください。

- [Go 公式サイト](https://golang.org/doc/install)

#### DevContainer

Visual Studio Code で Dev Container を使用することで、Docker コンテナ内での開発環境をすぐに構築できます。
必要なファイルは作成済みのため、以下の手順で開発環境を構築してください。

1. Visual Studio Code を開く
2. コマンドパレットを開く
3. `Remote-Containers: Reopen in Container` を選択
4. コンテナのビルドが始まり、完了後にコンテナ内で Visual Studio Code が開きます
5. 別途拡張機能や設定を追加したい場合は、コンテナ内で開いた Visual Studio Code で行ってください

詳細は Dev Container の使い方は以下のリンクを参考にしてください。

- [Developing inside a Container - Visual Studio Code](https://code.visualstudio.com/docs/devcontainers/containers)

#### Docker Compose

ローカルから Docker Compose を使用してテストを行う場合は Docker Desktop のインストールが必要です。
以下のリンクからインストール方法を確認してください。

- [Docker 公式サイト](https://docs.docker.com/get-docker/)

インストール後、以下のコマンドで Docker Compose が使えることを確認してください。

```bash
docker compose ps
```

なお Docker Desktop 以外の Docker 環境を使用する場合は、適宜インストール方法を調べてください。

### タスクランナーのインストール

テストの実行用に以下のタスクランナーを使用することができます。

- [Task](https://taskfile.dev/)

利用しなくともテストの実行は可能ですが、すぐにインストールできるためよろしければ導入ください。

#### Via Go Module

Go をローカルにインストール済みの方は以下のコマンドで導入可能です。

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

#### For Mac

```bash
brew install go-task/tap/go-task
```

#### For Linux

```bash
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin
```

#### For Windows

```bash
choco install go-task
```

## Usage

### 当リポジトリの使い方

当リポジトリは Go Tutorial の内容と連動したコーディング課題を提供しています。
ご自身の名前でブランチを切り、各課題を実装して当リポジトリに Push してください。
Push 後に走る CI が無事に通れば課題クリアです。

`./subject` ディレクトリ以下に課題が格納されています。
`chapterXX` というディレクトリがあり、その中に各課題とそのテストが格納されています。
chapter ごとの課題の詳細が README.md に記載されているので、そちらを参照して実装を行ってください。

::: info
`feature/subject-examples` というブランチに模範解答があります。
こちらは基本的に解説時に使用するため、自分で実装する前に見るのはおすすめしません。
:::

::: warning
PR を出したり、 `main` および `feature/subject-examples` に直接コミットするのは禁止です。
:::

### テストの実行

Golang は標準でテスト実行用のコマンドが用意されており、以下のコマンドでテストが実行できます。

```bash
# すべてのテストを実行する場合
go test ./...

# 特定のパッケージのテストを実行する場合
go test ./<package-name>
```

また、タスクランナーにテスト実行用のタスクを定義済みです。

以下から適したタスクを走らせてテストを実行してください。

#### ローカルの Go バイナリを使用して or Container 内でテストを実行する場合

```bash
task test
```

特定のパッケージのみテストしたい場合は以下のように `--` の後にパッケージ名を渡して実行してください。

```bash
task test -- <package-name>
```

#### ローカルから Compose を使ってテストする場合

初めて利用する場合はまず docker compose でビルドを行います。

```bash
docker compose build
```

ビルド済みの状態で以下のコマンドを実行すると Compose が立ち上がります。

```bash
docker compose up -v
```

Compose が立ち上がると `app` というサービスが立ち上がります。
この中が実行環境になっているので、以下のように `docker compose exec` コマンドを使ってテストやその他シェル操作を行ってください。

```bash
docker compose exec app task test -- chapter00
```

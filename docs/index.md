# Hello Golang

## Subjects

| Chapter    | Link                                                                                     | Workflow                                                                                 |
| :--------- | :--------------------------------------------------------------------------------------- | :--------------------------------------------------------------------------------------- |
| Chapter 00 | [Marp](./marp/installation.html)                                                         | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter00_test.yml) |
| Chapter 01 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter01/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter01_test.yml) |
| Chapter 02 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter02/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter02_test.yml) |
| Chapter 03 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter03/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter03_test.yml) |
| Chapter 04 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter04/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter04_test.yml) |
| Chapter 05 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter05/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter05_test.yml) |
| Chapter 06 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter06/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter06_test.yml) |
| Chapter 07 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter07/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter07_test.yml) |
| Chapter 08 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter08/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter08_test.yml) |
| Chapter 09 | [README](https://github.com/kurupeku/hello-golang/tree/main/subject/chapter09/README.md) | [Actions](https://github.com/kurupeku/hello-golang/actions/workflows/chapter09_test.yml) |

## Get Started

### Go のローカルインストール

#### Mac

##### インストーラー経由でインストール（推奨）

[こちら](https://go.dev/doc/install) からダウンロードして実行

##### Homebrew を使用してインストール（バージョンが古い場合あり）

1. `brew install go`
1. `go version` でバージョンが表示されれば OK

#### Windows

[こちら](https://go.dev/doc/install) からインストーラーをダウンロードして実行

#### Linux

1. `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz`
1. `vi ~/.bashrc` で `export PATH=$PATH:/usr/local/go/bin` を追記
1. `source ~/.bashrc`
1. `go version` でバージョンが表示されれば OK

### タスクランナーのインストール

テストの実行用に以下のタスクランナーを使用する想定です。

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

### ローカルから Docker コンテナ内の Shell を操作する

タスクランナーに以下のコマンドを定義しています。

```bash
task docker-bash # bashで接続

task docker-zsh  # zsh で接続
```

お好きな方の Shell を選んでタスクを実行すると、 Container が立ち上がって内部の Shell に接続された状態になります。

`--rm` フラグをつけているのでコンテナは都度破棄されます。

### テストの実行

Golang は標準でテスト実行用のコマンドが用意されており、以下のコマンドで実行することができます。

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

#### ローカルから Docker を利用してテストを実行する場合

初めて利用する場合はまず DockerImage をビルドします。

```bash
task docker-build
```

DockerImage がビルド済みの状態で以下のコマンドを実行すると DockerContainer が立ち上がり、 Container 内でテストコマンドを実行することができます

```bash
task docker-test

# ローカルと同じようにパッケージ指定も可能です
task docker-test -- <package-name>
```

上記 2 つを同時に実行することもできます。

```bash
task docker-build-test
```

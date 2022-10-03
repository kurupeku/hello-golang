# hello-golang

Subjects for learning golang on the "A Tour of Go"

## Pages

- https://kurupeku.github.io/hello-golang/

## Get Started

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

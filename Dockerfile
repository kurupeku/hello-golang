FROM golang:bullseye

LABEL maintainer "kurupeku <y.u.kurupeku@gmail.com>"

ENV GOENV dev
ENV GOPATH /go
ENV GOBIN ${GOPATH}/bin
ENV ROOT ${GOPATH}/src/app
ENV CGO_ENABLED 0

ENV PACKAGES zsh neovim git curl jq ripgrep

WORKDIR ${ROOT}

RUN sed -i 's@archive.ubuntu.com@ftp.jaist.ac.jp/pub/Linux@g' /etc/apt/sources.list \
  && apt-get update \
  && apt-get install -y --no-install-recommends ${PACKAGES} \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && chsh -s /bin/zsh \
  && go install github.com/go-task/task/v3/cmd/task@latest \
  && go install golang.org/x/tools/gopls@latest \
  && go install honnef.co/go/tools/cmd/staticcheck@latest \
  && go install golang.org/x/tools/cmd/goimports@latest \
  && go install mvdan.cc/gofumpt@latest \
  && go install github.com/cweill/gotests/gotests@latest \
  && go install github.com/fatih/gomodifytags@latest \
  && go install github.com/josharian/impl@latest \
  && go install github.com/haya14busa/goplay/cmd/goplay@latest \
  && go install github.com/jesseduffield/lazygit@latest \
  && go install mvdan.cc/sh/v3/cmd/shfmt@latest
COPY ./ ./
RUN go mod download

EXPOSE 3000

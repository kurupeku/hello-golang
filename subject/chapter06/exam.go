package chapter06

import (
	"strconv"

	"github.com/kurupeku/hello-golang/helper"

	"strings"
)

type AhoNumber int

func (n AhoNumber) String() string {
	return strconv.Itoa(int(n))
}

func (n AhoNumber) aho() string {
	return helper.AhoString(int(n))
}

func (n AhoNumber) Call() string {
	// TODO: 実装
	// nを3で割った余りが0 or nを文字列化したものに3が含まれる場合は、nをアホにしたものを返す
	if int(n)%3 == 0 || strings.Contains(n.String(), "3") {
		return n.aho()
	}
	// 上記に当てはまらない場合は、nを文字列化したものを返す
	return n.String()
}

func Nabeatsu(n int) []string {
	// TODO: 実装
	// 戻り値用のsliceを定義
	var result []string

	// 1～nまで(nを含む)ループ
	for i := 1; i <= n; i++ {
		// iはint型なので、AhoNumberにキャストしahoに格納
		aho := AhoNumber(i)
		// 上記でキャストした数値(aho)を元にCallを呼び出し、結果をslice(result)へ追加していく
		result = append(result, aho.Call())
	}
	// resultを返す
	return result
}

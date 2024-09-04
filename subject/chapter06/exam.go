package chapter06

import (
	"strconv"

	"github.com/kurupeku/hello-golang/helper"
)

type AhoNumber int

func (n AhoNumber) String() string {
	return strconv.Itoa(int(n))
}

func (n AhoNumber) aho() string {
	return helper.AhoString(int(n))
}

// AhoNumber 自身の値に応じて、数値をそのまま文字列で返すか、アホになるかを判定して返してください
func (n AhoNumber) Call() string {
	// TODO: 実装
	return ""
}

// 1 から n までの数をアホになるかどうかを判定して、すべての結果を文字列のスライスで返してください
func Nabeatsu(n int) []string {
	// TODO: 実装
	return []string{}
}

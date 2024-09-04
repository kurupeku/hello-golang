package chapter06

import (
	"strconv"
	"strings"

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
	if n%3 != 0 && !strings.ContainsAny(n.String(), "3") {
		return n.String()
	}

	return n.aho()
}

// 1 から n までの数をアホになるかどうかを判定して、すべての結果を文字列のスライスで返してください
func Nabeatsu(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		a := AhoNumber(i)
		res = append(res, a.Call())
	}

	return res
}

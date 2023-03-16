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

func (n AhoNumber) Call() string {
	// TODO: 実装
	s := n.String()
	if n%3 == 0 || strings.Contains(s, "3") {
		return n.aho()
	}
	return s
}

func Nabeatsu(n int) []string {
	// TODO: 実装
	var ret []string
	for i := 1; i <= n; i++ {
		ahoNum := AhoNumber(i)
		s := ahoNum.Call()
		ret = append(ret, s)
	}
	return ret
}

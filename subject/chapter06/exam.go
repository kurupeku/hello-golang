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
	if n%3 == 0 || strings.Contains(n.String(), "3") {
		return n.aho()
	}
	return n.String()
}

func Nabeatsu(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = AhoNumber(i + 1).Call()
	}
	return ret
}

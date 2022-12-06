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
	a := ""
	if s := n.String(); int(n)%3 == 0 || strings.Contains(s, "3") {
		a = n.aho()
	} else {
		a = s
	}
	return a
}

func Nabeatsu(n int) []string {
	aho := make([]string, n)
	for i := 1; i <= n; i++ {
		a := AhoNumber(i)
		aho[i-1] = a.Call()
	}
	return aho
}

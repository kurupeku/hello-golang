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
	if int(n)%3 == 0 {
		return n.aho()
	} else if strings.Contains(n.String(), "3") {
		return n.aho()
	}
	return n.String()
}

func Nabeatsu(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		result[i-1] = AhoNumber(i).Call()
	}
	return result
}

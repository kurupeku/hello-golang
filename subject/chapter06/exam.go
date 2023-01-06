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
	ahoitoa := n.String()

	if strings.Contains(ahoitoa, "3") || n%3 == 0 {
		return n.aho()
	}

	return ahoitoa
}

func Nabeatsu(n int) []string {
	ahoStrings := make([]string, n)

	for i := 1; i <= n; i++ {
		ahoStrings[i-1] = AhoNumber(i).Call()
	}

	return ahoStrings
}

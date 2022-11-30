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

	if n%3 == 0 {
		return n.aho()
	}

	if strings.Contains(n.String(), "3") {
		return n.aho()
	}

	return n.String()
}

func Nabeatsu(n int) []string {
	// TODO: 実装

	answer := []string{}

	for i := 1; i <= n; i++ {
		answer = append(answer, AhoNumber(i).Call())
	}

	return answer
}

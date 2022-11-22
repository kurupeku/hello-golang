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
	if n%3 == 0 {
		return n.aho()
	} else if strings.Index(n.String(), "3") != -1 {
		return n.aho()
	}
	return n.String()
}

func Nabeatsu(n int) []string {
	// TODO: 実装
	array := []string{}
	for i := 1; i <= n; i++ {
		array = append(array, AhoNumber(i).Call())
	}
	return array
}

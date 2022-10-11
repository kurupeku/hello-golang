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

func (n AhoNumber) Call() string {
	// TODO: 実装
	return ""
}

func Nabeatsu(n int) []string {
	// TODO: 実装
	return []string{}
}

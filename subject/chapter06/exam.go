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
	if int(n)%3 == 0 {
		return n.aho()
	} else if strings.Contains(n.String(), "3") {
		return n.aho()
	} else {
		return n.String()
	}
}

func Nabeatsu(n int) []string {
	// TODO: 実装
	var x []string
	for i := 1; i <= n; i++ {
		aho := AhoNumber(i)
		x = append(x, aho.Call())
		//	fmt.Println(x[i])
	}
	return x
}

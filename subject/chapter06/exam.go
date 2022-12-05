package chapter06

import (
	"fmt"
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
	fmt.Printf("Callのnは %d\n", n)

	switch {
	// 3の倍数
	case n%3 == 0:
		return n.aho()
	// 3がつく時
	case strings.Contains(n.String(), "3"):
		return n.aho()
	// その他
	default:
		return n.String()
	}

}

func Nabeatsu(n int) []string {
	var a []string

	for i := 1; i <= n; i++ {
		a = append(a, AhoNumber(i).Call())
	}

	return a
}

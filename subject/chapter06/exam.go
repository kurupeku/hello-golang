package chapter06

import (
	"fmt"
	"strconv"
	"strings"
)

type AhoNumber int

func (n AhoNumber) String() string {
	return strconv.Itoa(int(n))
}

func (n AhoNumber) aho() string {
	return fmt.Sprintf("%d aho", n)
}

func (n AhoNumber) Call() string {
	if n%3 != 0 && !strings.ContainsAny(n.String(), "3") {
		return n.String()
	}

	return n.aho()
}

func Nabeatsu(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		a := AhoNumber(i)
		res = append(res, a.Call())
	}

	return res
}

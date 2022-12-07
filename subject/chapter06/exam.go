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
	if (n % 3 == 0) {
		return n.aho();
	}

	nChar := n.String()
	if (strings.Contains(nChar, "3")) {
		return n.aho();
	}

	return nChar;
}

func Nabeatsu(n int) []string {
	var ans []string
	var aho AhoNumber
	for i := 1; i <= n; i++ {
		aho = AhoNumber(i)
		ans = append(ans, aho.Call())
	}
	return ans;
}

package chapter06

import (
	"strconv"
	"strings"

	"github.com/kurupeku/hello-golang/helper"
)

type AhoNumber int

const nabeNum int = 3

func (n AhoNumber) String() string {
	return strconv.Itoa(int(n))
}

func (n AhoNumber) aho() string {
	return helper.AhoString(int(n))
}

func (n AhoNumber) Call() string {

	if (int(n) % nabeNum) == 0 {
		return n.aho()
	} else if strings.Contains(n.String(), strconv.Itoa(nabeNum)) {
		return n.aho()
	} else {
		return n.String()
	}
}

func Nabeatsu(n int) []string {

	numAry := make([]string, n)

	for countNum := AhoNumber(1); int(countNum) <= n; countNum++ {
		numAry[countNum-1] = countNum.Call()
	}

	// TODO: 実装
	return numAry
}

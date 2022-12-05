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
	Response := n.String()

	if strings.Contains(Response, "3") {
		Response = helper.AhoString(int(n))
		return Response
	}

	if n%3 == 0 {
		Response = helper.AhoString(int(n))
		return Response
	}
	return Response
}

func Nabeatsu(n int) []string {
	var answer []string = nil
	for i := 1; i <= n; i++ {
		m := AhoNumber(i)
		var pre_answer string = AhoNumber.Call(m)
		answer = append(answer, pre_answer)
	}
	return answer
}

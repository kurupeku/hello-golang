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
	arg := AhoNumber(n)
	arg_str := arg.String()
	multiple_three := n % 3
	switch {
	case strings.Contains(arg_str, "3"):
		aho := arg.aho()
		return aho
	case multiple_three == 0:
		aho := arg.aho()
		return aho
	default:
		return arg_str
	}
}

func Nabeatsu(n int) []string {
	// TODO: 実装
	str_array := make([]string, 0)
	for i := 1; i <= n; i++ {
		arg := AhoNumber(i)
		arg_str := arg.String()
		multiple_three := i % 3
		switch {
		case strings.Contains(arg_str, "3"):
			saaaan := arg.aho()
			str_array = append(str_array, saaaan)
		case multiple_three == 0:
			saaaan := arg.aho()
			str_array = append(str_array, saaaan)
		default:
			aho := AhoNumber(i)
			str_array = append(str_array, aho.String())
		}

	}
	return str_array
}

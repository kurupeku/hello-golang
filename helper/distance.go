package helper

import (
	"regexp"
	"strconv"
)

type distance string

func (d distance) value() float32 {
	reg := regexp.MustCompile(`^[0-9.]+`)
	i, err := strconv.ParseFloat(reg.FindString(string(d)), 32)
	if err != nil {
		panic(err)
	}

	return float32(i)
}

func (d distance) measure() float32 {
	reg := regexp.MustCompile(`\D+$`)
	switch reg.FindString(string(d)) {
	case "km":
		return 1000
	case "cm":
		return 0.01
	case "mm":
		return 0.001
	default:
		return 1
	}
}

// 距離を表す文字列を渡すとメートルに直した上で数値として返却する関数
//
// @Params dist string 距離を表す文字列
//
//	数値のみの文字列はメートルとして扱われる
//	e.g.) "1200m" / "1.85km" / "350000cm"
//
// @Return int メートルに直された距離
//
//	e.g.) 1.82km => 1820
func ParseDistance(dist string) int {
	d := distance(dist)

	return int(d.value() * d.measure())
}

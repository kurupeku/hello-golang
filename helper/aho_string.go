package helper

import (
	"fmt"
	"strings"
)

var handredsDigit = map[string]string{
	"0": "",
	"1": "ひゃく",
	"2": "にひゃく",
	"3": "さんびゃく",
	"4": "よんひゃく",
	"5": "ごひゃく",
	"6": "ろっぴゃく",
	"7": "ななひゃく",
	"8": "はっぴゃく",
	"9": "きゅうひゃく",
}

var tensDigit = map[string]string{
	"0": "",
	"1": "じゅう",
	"2": "にじゅう",
	"3": "さんじゅう",
	"4": "よんじゅう",
	"5": "ごじゅう",
	"6": "ろくじゅう",
	"7": "ななじゅう",
	"8": "はちじゅう",
	"9": "きゅうじゅう",
}

var onesDigit = map[string]string{
	"0": "",
	"1": "いぃち",
	"2": "にぃ",
	"3": "さぁん",
	"4": "しぃ",
	"5": "ごぉ",
	"6": "ろぉく",
	"7": "しぃち",
	"8": "はぁち",
	"9": "きゅう",
}

func AhoString(i int) string {
	zv := fmt.Sprintf("%03d", i)
	za := strings.Split(zv, "")
	return handredsDigit[za[0]] + tensDigit[za[1]] + onesDigit[za[2]]
}

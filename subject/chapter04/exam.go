package chapter04

import (
	"fmt"
)

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	tmp := 0

	switch {
	case charge < card.Point:
		card.Point -= charge
		fmt.Println(card)
		return true
	case charge <= card.Balance+card.Point:
		tmp = charge - card.Point
		card.Point = 0
		card.Balance -= tmp
		fmt.Println(card)
		return true
	}

	return false
}

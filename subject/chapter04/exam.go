package chapter04

import (
// "fmt"
)

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装

	/*
		fmt.Printf("乗車料金:%d\n", charge)
		fmt.Printf("card.Balance:%d\n", card.Balance)
		fmt.Printf("card.Point:%d\n", card.Point)
	*/

	var total int
	var tmp int

	total = card.Balance + card.Point

	// そもそもICカード残高より乗車料金が高いかどうか
	// 残高以内
	if total-charge >= 0 {
		// カードポイントに余りがある場合
		if card.Point > 0 {
			tmp = card.Point - charge
			// fmt.Printf("tmp: %d\n", tmp)
			if tmp <= 0 {
				// print("1\n")
				card.Point = card.Point - charge
				card.Point = card.Point - tmp
				card.Balance = card.Balance + tmp
				// fmt.Printf("end card.Point: %d\n", card.Point)
				// fmt.Printf("end card.Balance: %d\n", card.Balance)
				return true
			} else {
				// print("2\n")
				card.Point = card.Point - charge
				return true
			}
			// カードポイント残高0の場合
		} else {
			card.Balance = card.Balance - charge
			return true
		}
	} else {
		// そもそも初めから残高不足だからfalse
		return false
	}
}

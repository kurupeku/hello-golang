package chapter09

import "fmt"

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// TODO: 実装

	// Ticket or Card
	var i interface{} = charger

	switch v := i.(type) {
	case *Ticket:
		fmt.Printf("Type is Ticket = %T", v)
		// 距離を確認

		// 利用料金を算出
	case *Card:
		fmt.Printf("Type is Card = %T", v)
	default:
		fmt.Printf("Type is Unknown = %T", v)
	}

	// 利用料金(charge)を算出

	return false, nil
}

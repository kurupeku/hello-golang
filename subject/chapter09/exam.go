package chapter09

import (
	"errors"
	"fmt"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// TODO: 実装

	// 料金
	var price int

	// chargeのタイプがチケットかカードか
	// fromとtoからかかる料金を算出
	fare := Fare{from, to}
	switch charger.(type) {
	case *Ticket:
		price = fare.TicketCharge()
	case *Card:
		price = fare.CardCharge()
	}

	fmt.Printf("price : %d\n", price)

	// 利用料金が0の場合はエラーを返す
	if price == 0 {
		return false, errors.New("利用料金が0円です。")
	}

	// 残額（チケット or カード）が利用料金より多ければ支払いはできる。
	// 少なければ支払いはできない。当たり前。
	if charger.Amount() < price {
		return false, nil
	}

	charger.Use(price)

	return true, nil
}

package chapter09

import (
	"errors"
	"fmt"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// TODO: 実装

	station := Fare{from, to}
	price := 0
	switch charger.(type) {
	case *Ticket:
		price = station.TicketCharge()
		fmt.Println(price)
	case *Card:
		price = station.CardCharge()
		fmt.Println(price)
	}

	if from == to {
		return false, errors.New("error:利用金額が0です。")
	} else if charger.Amount()-price < 0 {
		return false, nil
	} else {
		charger.Use(price)
		return true, nil
	}

}

package chapter09

import "errors"

var errAmount0 = errors.New("amount0 error") // エラーメッセージはこうまとめるお作法？

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	var charge int
	fare := Fare{From: from, To: to}

	if fare.distance() == 0 {
		return false, errAmount0
	}

	switch charger.(type) {
	case *Ticket:
		charge = fare.TicketCharge()
	case *Card:
		charge = fare.CardCharge()
	}

	result := charger.Amount() > charge
	if result {
		charger.Use(charge)
	}

	return result, nil
}

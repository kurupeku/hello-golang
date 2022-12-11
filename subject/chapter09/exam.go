package chapter09

import (
	"errors"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	fare := Fare{From: from, To: to}
	var amount int
	switch charger.(type) {
	case *Ticket:
		amount = fare.TicketCharge()
	case *Card:
		amount = fare.CardCharge()
	default:
		return false, errors.New("invalid Charger type")
	}

	if amount == 0 {
		return false, errors.New("invalid input No Fare")
	}

	if charger.Amount() < amount {
		return false, nil
	}

	charger.Use(amount)
	return true, nil
}

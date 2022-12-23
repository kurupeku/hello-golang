package chapter09

import (
	"errors"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// TODO: 実装

	fare := Fare{from, to}
	var charge int

	switch charger.(type) {
	case *Ticket:
		if fare.TicketCharge() == 0 {
			return false, errors.New("TicketCharge is 0")
		}
		charge = fare.TicketCharge()

	case *Card:
		if fare.CardCharge() == 0 {
			return false, errors.New("CardCharge is 0")
		}
		charge = fare.CardCharge()

	default:
		return false, errors.New("unexpected error")
	}

	if charger.Amount() < charge {
		return false, nil
	}

	charger.Use(charge)

	return true, nil
}

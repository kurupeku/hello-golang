package chapter09

import (
	"errors"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	if from == to {
		error := errors.New("same station")
		return false, error
	}

	a := charger.Amount()
	if a == 0 {
		error := errors.New("amount is zero")
		return false, error
	}

	f := &Fare{from, to}

	switch charger.(type) {
	case *Ticket:
		fare := f.TicketCharge()
		if fare > a {
			return false, nil
		}

		charger.Use(fare)
	case *Card:
		fare := f.CardCharge()
		if fare > a {
			return false, nil
		}

		charger.Use(fare)
	default:
		break
	}

	return true, nil
}

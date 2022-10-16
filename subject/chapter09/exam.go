package chapter09

import (
	"errors"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	var charge int
	f := Fare{from, to}
	switch charger.(type) {
	case *Ticket:
		charge = f.TicketCharge()
	case *Card:
		charge = f.CardCharge()
	}

	if charge == 0 {
		return false, errors.New("could not calculate charge")
	}

	if charge > charger.Amount() {
		return false, nil
	}

	charger.Use(charge)
	return true, nil
}

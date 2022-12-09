package chapter09

import (
	"errors"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// TODO: 実装
	if from == to {
		return false, errors.New("dame")
	}
	v := Fare{from, to}
	var charge int
	switch charger.(type) {
	case *Card:
		charge = v.CardCharge()
	case *Ticket:
		charge = v.TicketCharge()
	}
	if charger.Amount() < charge {
		return false, nil
	}
	charger.Use(charge)
	return true, nil
}

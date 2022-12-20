package chapter09

import (
	"errors"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// TODO: 実装
	var f Fare
	f.From = from
	f.To = to
	var charge int

	if f.distance() == 0 {
		return false, errors.New("zero yen")
	}

	if _, ok := charger.(*Card); ok {
		charge = f.CardCharge()
	} else {
		charge = f.TicketCharge()
	}

	if charger.Amount() < charge {
		return false, nil
	}

	charger.Use(charge)
	return true, nil
}

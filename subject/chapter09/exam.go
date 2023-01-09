package chapter09

import (
	"errors"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	fare := Fare{
		From: from,
		To:   to,
	}
	var fine int
	switch charger.(type) {
	case *Ticket:
		fine = fare.TicketCharge()
	case *Card:
		fine = fare.CardCharge()
	default:
		return false, errors.New("エラー文")
	}

	if fine == 0 {
		return false, errors.New("エラー文")
	}

	if charger.Amount() < fine {
		return false, nil
	}
	charger.Use(fine)
	return true, nil
}

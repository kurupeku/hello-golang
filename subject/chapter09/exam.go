package chapter09

import "errors"

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	fare := Fare{
		From: from,
		To:   to,
	}

	if fare.From == fare.To {
		return false, errors.New("[ERROR] 発着駅が同一です")
	}

	var pay int
	switch charger.(type) {
	case *Ticket:
		pay = fare.TicketCharge()
	case *Card:
		pay = fare.CardCharge()
	default:
		return false, nil
	}

	if charger.Amount() >= pay {
		charger.Use(pay)
		return true, nil
	} else {
		return false, nil
	}
}

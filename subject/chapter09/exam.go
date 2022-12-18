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
	if _, isTicket := charger.(*Ticket); isTicket {
		pay = fare.TicketCharge()
	} else if _, isCard := charger.(*Card); isCard {
		pay = fare.CardCharge()
	}

	if charger.Amount() >= pay {
		charger.Use(pay)
		return true, nil
	} else {
		return false, nil
	}
}

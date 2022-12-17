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

	if _, isTicket := charger.(*Ticket); isTicket {
		pay := fare.TicketCharge()
		if charger.Amount() >= pay {
			charger.Use(pay)
			return true, nil
		} else {
			return false, nil
		}
	} else if _, isCard := charger.(*Card); isCard {
		pay := fare.CardCharge()
		if charger.Amount() >= pay {
			charger.Use(pay)
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, nil
}

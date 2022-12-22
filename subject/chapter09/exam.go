package chapter09

import "errors"

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	if from == to {
		return false, errors.New("エラー")		
	}

	var canPass bool
	fare := Fare{from, to}
	switch charger := charger.(type) {
		case *Ticket:
			charger.Use(fare.TicketCharge())
			canPass = charger.Used
		case *Card:
			beforeAmount := charger.Amount()
			charger.Use(fare.CardCharge())
			canPass = beforeAmount != charger.Amount()
	}
	return canPass, nil;
}
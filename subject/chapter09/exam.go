package chapter09

import "errors"

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	f := new(Fare)
	f.From = from
	f.To = to

	if f.From == f.To {
		return false, errors.New("乗車料金は0円です")
	}

	switch charger.(type) {
	case *Ticket:
		result := charger.Use(f.TicketCharge())
		if result {
			return true, nil
		}
		return false, nil
	case *Card:
		result := charger.Use(f.CardCharge())
		if result {
			return true, nil
		}
		return false, nil
	}

	return false, nil
}

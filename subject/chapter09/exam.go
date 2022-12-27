package chapter09

import (
	"fmt"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	if from == to {
		return false, fmt.Errorf("error: %s", "乗車駅と降車駅が一緒だよ")
	}

	f := Fare{from, to}
	if f.distance() == 0 {
		return false, fmt.Errorf("error: %s", "距離が0だよ")
	}

	var charge int
	if _, ok := charger.(*Ticket); ok {
		charge = f.TicketCharge()
	}
	if _, ok := charger.(*Card); ok {
		charge = f.CardCharge()
	}

	if charger.Amount() < charge {
		return false, nil
	}
	charger.Use(charge)

	return true, nil
}

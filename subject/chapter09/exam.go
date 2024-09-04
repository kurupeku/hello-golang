package chapter09

import (
	"errors"
)

// 出発駅から到着駅までの料金を計算し、渡された切符 or IC カードで支払いを行う関数を実装してください
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

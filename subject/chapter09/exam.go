package chapter09

import (
	"errors"
	// "fmt"
)

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// Ticket or Card
	var i interface{} = charger
	f := Fare{from, to}

	switch i.(type) {
	case *Ticket:
		/*
			fmt.Printf("Type is Ticket = %T \n", v)
			fmt.Printf("TicketCharge Value = %v \n", f.TicketCharge())
			fmt.Printf("Amount Value = %v \n", charger.Amount())
		*/

		switch {
		//乗車不可
		case charger.Amount() < f.TicketCharge():
			return false, nil
		//利用料金ゼロ円
		case f.TicketCharge() == 0:
			return false, errors.New("zero")
		//乗車可能
		default:
			charger.Use(f.TicketCharge())
			return true, nil
		}

		// 利用料金を算出
	case *Card:
		/*
			fmt.Printf("Type is Card = %T \n", v)
			fmt.Printf("CardCharge Value = %v \n", f.CardCharge())
			fmt.Printf("Amount Value = %v \n", charger.Amount())
		*/

		switch {
		//乗車不可
		case charger.Amount() < f.CardCharge():
			return false, nil
		//利用料金ゼロ円
		case f.CardCharge() == 0:
			return false, errors.New("zero")
		//乗車可能
		default:
			charger.Use(f.CardCharge())
			return true, nil
		}

	default:
		return false, errors.New("type is unknown")
		//fmt.Printf("Type is Unknown = %T", v)
	}

}

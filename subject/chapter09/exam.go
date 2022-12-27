package chapter09

import "fmt"

type MyError struct {
	Message string
}

func (err *MyError) Error() string {
	return err.Message
}

func err() error {
	return &MyError{"残高不足"}
}

func Kaisatsu(from, to string, charger Charger) (bool, error) {
	// TODO: 実装

	stations := Fare{from, to}
	price := 0

	switch charger.(type) {
	case *Ticket:
		price = stations.TicketCharge()
		fmt.Println(price)
	case *Card:
		price = stations.CardCharge()
		fmt.Println(price)
	default:
		return false, nil
	}

	if price == 0 {
		return false, err()

	} else if charger.Amount()-price < 0 {
		return false, nil
	} else {
		charger.Use(price)
		return true, nil
	}
}

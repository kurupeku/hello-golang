package chapter09

type Charger interface {
	Amount() int
	Use(charge int)
}

// Charger を満たすようにメソッドを追加する
type Ticket struct {
	Price int
	Used  bool
}

// Charger を満たすようにメソッドを追加する
type Card struct {
	Balance int
	Point   int
}

func (ticket *Ticket) Amount() int {
	return ticket.Price
}

func (card *Card) Amount() int {
	return card.Balance + card.Point
}

func (ticket *Ticket) Use(charge int) {
	ticket.Used = true
}

func (card *Card) Use(charge int) {
	if charge > card.Point {
		charge -= card.Point
		card.Balance -= charge
		card.Point = 0
	} else {
		card.Point -= charge
	}
}

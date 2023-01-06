package chapter09

type Charger interface {
	Amount() int
	Use(charge int)
}

type Ticket struct {
	Price int
	Used  bool
}

type Card struct {
	Balance int
	Point   int
}

func (card *Card) Amount() int {
	return card.Balance + card.Point
}

func (card *Card) Use(charge int) {
	if card.Point > charge {
		card.Point -= charge
	} else {
		card.Balance -= (charge - card.Point)
		card.Point = 0
	}
}

func (ticket *Ticket) Amount() int {
	return ticket.Price
}

func (ticket *Ticket) Use(charge int) {
	ticket.Used = true
}

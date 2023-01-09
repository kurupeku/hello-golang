package chapter09

type Charger interface {
	Amount() int
	Use(charge int)
}

type Ticket struct {
	Price int
	Used  bool
}

func (t *Ticket) Amount() int {
	return t.Price
}

func (t *Ticket) Use(charger int) {
	if t.Price >= charger {
		t.Used = true
	}
}

type Card struct {
	Balance int
	Point   int
}

func (c *Card) Amount() int {
	return c.Balance
}

func (c *Card) Use(charger int) {
	if c.Point >= charger {
		c.Point -= charger
	} else {
		charger -= c.Point
		c.Point = 0
		c.Balance -= charger
	}
}

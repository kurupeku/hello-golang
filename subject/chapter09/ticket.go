package chapter09

type Charger interface {
	Amount() int
	Use(charge int) bool
}

type Ticket struct {
	Price int
	Used  bool
}

func (t *Ticket) Amount() int {
	return t.Price
}

func (t *Ticket) Use(charge int) bool {
	if t.Used {
		return false
	}

	if t.Amount() < charge {
		return false
	}
	t.Used = true
	return true
}

type Card struct {
	Balance int
	Point   int
}

func (c *Card) Amount() int {
	return c.Balance + c.Point
}

func (c *Card) Use(charge int) bool {
	if (c.Balance + c.Point) < charge {
		return false
	}

	if charge <= c.Point {
		c.Point = c.Point - charge
		return true
	}

	c.Balance = c.Balance - (charge - c.Point)
	c.Point = 0
	return true
}

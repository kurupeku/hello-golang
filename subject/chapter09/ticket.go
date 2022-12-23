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

func (c *Ticket) Amount() int {
	return c.Price
}

func (c *Ticket) Use(charge int) {
	if c.Price > charge {
		c.Used = true
	} else {
		c.Used = false
	}
}

func (c *Card) Amount() int {
	return c.Balance + c.Point
}

func (c *Card) Use(charge int) {

	if c.Point >= charge {
		c.Point -= charge
	} else {
		c.Balance = c.Balance + c.Point - charge
		c.Point = 0
	}
}

package chapter09

type Charger interface {
	Amount() int
	Use(charge int)
}

// Charger を満たすようにメソッドを定義し、その実装してください
type Ticket struct {
	Price int
	Used  bool
}

func (t *Ticket) Amount() int {
	return t.Price
}

func (t *Ticket) Use(int) {
	t.Used = true
}

// Charger を満たすようにメソッドを定義し、その実装してください
type Card struct {
	Balance int
	Point   int
}

func (c *Card) Amount() int {
	return c.Balance + c.Point
}

func (c *Card) Use(charge int) {
	if c.Point > charge {
		c.Point -= charge
		return
	}

	remain := charge - c.Point
	c.Balance -= remain
	c.Point = 0
}

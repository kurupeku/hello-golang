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

func (t *Ticket) Amount() int {
	return t.Price
}
func (t *Ticket) Use(charge int) {
	t.Used = true
}
func (c *Card) Amount() int {
	return c.Balance + c.Point
}
func (c *Card) Use(charge int) {
	tmp := c.Point - charge
	if tmp >= 0 {
		c.Point = tmp
	} else {
		c.Point = 0
		c.Balance += tmp
	}
}

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

func (t *Ticket) Amount() int {
	return t.Price
}

func (t *Ticket) Use(charge int) {
	t.Used = true
}

// Charger を満たすようにメソッドを追加する
type Card struct {
	Balance int
	Point   int
}

func (c *Card) Amount() int {
	return c.Balance + c.Point
}

func (c *Card) Use(charge int) {

	// 支払い残高
	paymentBalance := charge

	if c.Balance+c.Point < paymentBalance {
		return
	}

	// ポイントから優先的に支払う
	if c.Point >= charge {
		c.Point -= charge
	} else {
		paymentBalance -= c.Point
		c.Point = 0
		c.Balance -= paymentBalance
	}
}

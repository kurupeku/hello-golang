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

func (ticket Ticket) Amount() int {
	return ticket.Price
}

func (ticket *Ticket) Use(charge int) {
	if charge <= ticket.Price {
		ticket.Used = true
	}
}

// Charger を満たすようにメソッドを追加する
type Card struct {
	Balance int
	Point   int
}

func (card *Card) Amount() int {
	return card.Balance + card.Point
}

func (card *Card) Use(charge int) {
	// 足りない場合
	if charge > card.Point + card.Balance {
		return
	}
		
	// ポイントで払える場合
	if charge <= card.Point {
		card.Point -= charge
		return
	}

	// ポイントを使い切る
	charge -= card.Point
	card.Point = 0

	// 残りを現金で
	card.Balance -= charge
}

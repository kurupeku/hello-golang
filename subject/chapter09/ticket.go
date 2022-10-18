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

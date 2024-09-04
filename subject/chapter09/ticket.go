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

// Charger を満たすようにメソッドを定義し、その実装してください
type Card struct {
	Balance int
	Point   int
}

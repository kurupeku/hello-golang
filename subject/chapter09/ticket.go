package chapter09

import "fmt"

// Ticket と Card をまとめて改札に渡す際に用いる interface
type Charger interface {
	// 乗車料金として利用可能な額を int 型で返すメソッド
	Amount() int
	//乗車料金 charge を int 型で受け取り、自身を使用した際の処理を行うメソッド
	Use(charge int)
}

// Charger を満たすようにメソッドを追加する
type Ticket struct {
	Price int
	Used  bool
}

func (t Ticket) Amount() int {
	switch {
	case t.Used:
		return 0
	default:
		return t.Price
	}
}

func (t *Ticket) Use(charge int) {
	switch {
	case charge >= t.Price:
		t.Used = true
	}
}

// Charger を満たすようにメソッドを追加する
type Card struct {
	// >=0 の int をチャージ残高として保持するフィールド
	Balance int
	// >=0 の int を保有ポイントとして保持するフィールド
	Point int
}

func (c Card) Amount() int {
	return c.Balance + c.Point
}

func (c *Card) Use(charge int) {
	tmp := 0

	switch {
	case charge < c.Point:
		c.Point -= charge
		fmt.Println("*Card = ", c)

	case charge <= c.Balance+c.Point:
		tmp = charge - c.Point
		c.Point = 0
		c.Balance -= tmp
		fmt.Println("*Card = ", c)

		// 料金 > 所持金だった場合、Cardは消費されない
	}
}

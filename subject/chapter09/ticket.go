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

func (ticket *Ticket) Amount() int {
	// チケットが使用済み(true)の場合はチケットの金額を0にする
	if ticket.Used {
		ticket.Price = 0
	}
	// チケットの金額を返す
	return ticket.Price
}

func (card *Card) Amount() int {
	// チャージ残高とポイント残高を足した金額を返す
	return card.Balance + card.Point
}

func (ticket *Ticket) Use(charge int) {
	// チケットを使用済みにする
	ticket.Used = true
}

func (card *Card) Use(charge int) {
	// 残高+ポイントが乗車料金より多いかの判定
	if charge <= card.Balance+card.Point {
		// 乗車料金が0より大きい かつ ポイントが0より大きい場合にループ
		for charge > 0 && card.Point > 0 {
			// 乗車料金とポイントをループする度に1マイナスする
			charge--
			card.Point--
		}
		// ポイントから引けなかった分の乗車料金を残高からマイナスする（ポイントで足りた場合は0を引く）
		card.Balance -= charge
	}
}

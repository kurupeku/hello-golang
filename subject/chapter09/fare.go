package chapter09

type Fare struct {
	From string
	To   string
}

// 内回りの距離を返すように実装してください
func (f *Fare) innerDistance() int {
	// TODO: 実装
	return 0
}

// 外回りの距離を返すように実装してください
func (f *Fare) outerDistance() int {
	// TODO: 実装
	return 0
}

// 内回りと外回りの距離を比較し、短い方の距離を返すように実装してください
func (f *Fare) distance() int {
	// TODO: 実装
	return 0
}

// 切符の料金を返すように実装してください
func (f *Fare) TicketCharge() int {
	// TODO: 実装
	return 0
}

// IC カードの料金を返すように実装してください
func (f *Fare) CardCharge() int {
	// TODO: 実装
	return 0
}

package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// TODO: 内回りの料金を返すように実装
	return 0
}

func (f *Fare) outerDistance() int {
	// TODO: 外回りの距離を返すように実装
	return 0
}

func (f *Fare) distance() int {
	// 内回りと外回りの距離を比較し、短い方を返すように実装
	return 0
}

func (f *Fare) TicketCharge() int {
	// TODO: 切符の料金を返すように実装
	return 0
}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装
	return 0
}

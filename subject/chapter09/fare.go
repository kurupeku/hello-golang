package chapter09

import (
	"fmt"

	"github.com/kurupeku/hello-golang/helper"
)

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// TODO: 内回りの距離を返すように実装

	// 走行距離
	var mileage int

	// 現在の駅
	currentStation := f.From

	if currentStation == f.To {
		return mileage
	}

	for {
		nextStation := helper.InnerNextStation(currentStation)
		mileage += helper.InnerLoopDistance(nextStation)
		currentStation = nextStation

		if currentStation == f.To {
			break
		}
	}
	return mileage
}

func (f *Fare) outerDistance() int {
	// TODO: 外回りの距離を返すように実装

	// 走行距離
	var mileage int

	// 現在の駅
	currentStation := f.From

	if currentStation == f.To {
		return mileage
	}

	for {
		nextStation := helper.OuterNextStation(currentStation)
		mileage += helper.OuterLoopDistance(nextStation)
		currentStation = nextStation

		if nextStation == f.To {
			break
		}
	}
	return mileage
}

func (f *Fare) distance() int {
	// 内回りと外回りの距離を比較し、短い方を返すように実装

	innerMileage := f.innerDistance()
	outerMileage := f.outerDistance()

	if innerMileage > outerMileage {
		return outerMileage
	} else {
		return innerMileage
	}
}

func (f *Fare) TicketCharge() int {
	// TODO: 切符の料金を返すように実装

	// 料金
	var price int

	mileage := f.distance()

	fmt.Printf("mileage : %d\n", mileage)

	if mileage <= 0 {
		price = 0
	} else if mileage <= 3999 {
		price = 140
	} else if mileage <= 6999 {
		price = 160
	} else if mileage <= 10999 {
		price = 170
	} else if mileage <= 15999 {
		price = 200
	} else if mileage <= 20999 {
		price = 270
	} else if mileage <= 25999 {
		price = 350
	} else if mileage <= 30999 {
		price = 420
	} else {
		price = 490
	}

	return price
}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装

	// 料金
	var price int

	mileage := f.distance()

	fmt.Printf("mileage : %d\n", mileage)

	if mileage <= 0 {
		price = 0
	} else if mileage <= 3999 {
		price = 136
	} else if mileage <= 6999 {
		price = 157
	} else if mileage <= 10999 {
		price = 168
	} else if mileage <= 15999 {
		price = 198
	} else if mileage <= 20999 {
		price = 264
	} else if mileage <= 25999 {
		price = 341
	} else if mileage <= 30999 {
		price = 418
	} else {
		price = 484
	}

	return price
}

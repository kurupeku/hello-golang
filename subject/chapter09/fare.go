package chapter09

import "github.com/kurupeku/hello-golang/helper"

type fareTable struct {
	minRage     int
	maxRange    int
	ticketPrice int
	cardPrice   int
}

var fareTableList = []fareTable{
	{
		minRage:     0,
		maxRange:    3999,
		ticketPrice: 140,
		cardPrice:   136,
	},
	{
		minRage:     4000,
		maxRange:    6999,
		ticketPrice: 160,
		cardPrice:   157,
	},
	{
		minRage:     7000,
		maxRange:    10999,
		ticketPrice: 170,
		cardPrice:   168,
	},
	{
		minRage:     11000,
		maxRange:    15999,
		ticketPrice: 200,
		cardPrice:   198,
	},
	{
		minRage:     16000,
		maxRange:    20999,
		ticketPrice: 270,
		cardPrice:   264,
	},
	{
		minRage:     21000,
		maxRange:    25999,
		ticketPrice: 350,
		cardPrice:   341,
	},
	{
		minRage:     26000,
		maxRange:    30999,
		ticketPrice: 420,
		cardPrice:   418,
	},
	{
		minRage:     31000,
		maxRange:    -1,
		ticketPrice: 490,
		cardPrice:   484,
	},
}

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	var dist int
	for cs := helper.InnerNextStation(f.From); cs != f.From; cs = helper.InnerNextStation(cs) {
		dist += helper.InnerLoopDistance(cs)
		if cs == f.To {
			return dist
		}
	}
	return 0
}

func (f *Fare) outerDistance() int {
	var dist int
	for cs := helper.OuterNextStation(f.From); cs != f.From; cs = helper.OuterNextStation(cs) {
		dist += helper.OuterLoopDistance(cs)
		if cs == f.To {
			return dist
		}
	}
	return 0
}

func (f *Fare) distance() int {
	if f.innerDistance() < f.outerDistance() {
		return f.innerDistance()
	} else {
		return f.outerDistance()
	}
}

func (f *Fare) TicketCharge() int {
	dist := f.distance()
	if dist == 0 {
		return 0
	}
	var price int
	for _, v := range fareTableList {
		if dist >= v.minRage && (dist <= v.maxRange || v.maxRange == -1) {
			price = v.ticketPrice
			break
		}
	}
	return price
}

func (f *Fare) CardCharge() int {
	dist := f.distance()
	if dist == 0 {
		return 0
	}
	var price int
	for _, v := range fareTableList {
		if dist >= v.minRage && (dist <= v.maxRange || v.maxRange == -1) {
			price = v.cardPrice
			break
		}
	}
	return price
}

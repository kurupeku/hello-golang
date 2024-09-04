package chapter09

type Fare struct {
	From string
	To   string
}

// 内回りの距離を返すように実装してください
func (f *Fare) innerDistance() int {
	if f.From == f.To {
		return 0
	}

	var dist int
	current := f.From
	for i := 0; i < 30; i++ {
		current = helper.InnerNextStation(current)
		dist += helper.InnerLoopDistance(current)
		if current == f.To {
			break
		}
	}

	return dist
}

// 外回りの距離を返すように実装してください
func (f *Fare) outerDistance() int {
	if f.From == f.To {
		return 0
	}

	var dist int
	current := f.From
	for i := 0; i < 30; i++ {
		current = helper.OuterNextStation(current)
		dist += helper.OuterLoopDistance(current)
		if current == f.To {
			break
		}
	}

	return dist
}

// 内回りと外回りの距離を比較し、短い方の距離を返すように実装してください
func (f *Fare) distance() int {
	id := f.innerDistance()
	od := f.outerDistance()

	if id > od {
		return od
	}

	return id
}

// 切符の料金を返すように実装してください
func (f *Fare) TicketCharge() int {
	dist := f.distance()
	switch {
	case dist == 0:
		return 0
	case dist < 4000:
		return 140
	case dist < 7000:
		return 160
	case dist < 11000:
		return 170
	case dist < 16000:
		return 200
	case dist < 21000:
		return 270
	case dist < 26000:
		return 350
	case dist < 31000:
		return 420
	default:
		return 490
	}
}

// IC カードの料金を返すように実装してください
func (f *Fare) CardCharge() int {
	dist := f.distance()
	switch {
	case dist == 0:
		return 0
	case dist < 4000:
		return 136
	case dist < 7000:
		return 157
	case dist < 11000:
		return 168
	case dist < 16000:
		return 198
	case dist < 21000:
		return 264
	case dist < 26000:
		return 341
	case dist < 31000:
		return 418
	default:
		return 484
	}
}

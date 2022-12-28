package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct {
	X int
	Y int
}

func (b *Basic) PricePer15Minutes() int {
	return 220
}

func (b *Basic) MaxPrice() int {
	return 4290
}

type Middle struct {
	X int
	Y int
}

func (m *Middle) PricePer15Minutes() int {
	return 330
}

func (m *Middle) MaxPrice() int {
	return 6490
}

type Premium struct {
	X int
	Y int
}

func (p *Premium) PricePer15Minutes() int {
	return 440
}

func (p *Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	price := 0
	count_6hour := minutes / 360
	amari_6hour := minutes % 360
	if count_6hour > 0 {
		price += car.MaxPrice() * count_6hour
	}
	count_15minute := amari_6hour / 15
	tuika_price := car.PricePer15Minutes() * (count_15minute + 1)
	if tuika_price > car.MaxPrice() {
		tuika_price = car.MaxPrice()
	}
	price += tuika_price

	return price
}

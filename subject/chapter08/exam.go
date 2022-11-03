package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

type Middle struct{}

type Premium struct{}

func Calc(car Car, minutes int) int {
	price := car.PricePer15Minutes()
	maxPrice := car.MaxPrice()

	d6h, m6h := divmod(minutes, 360)
	d15m, m15m := divmod(m6h, 15)

	if m15m != 0 {
		d15m++
	}

	var p1, p2 int
	p1 = d6h * maxPrice
	if p2 = d15m * price; p2 >= maxPrice {
		p2 = maxPrice
	}
	return p1 + p2
}

func divmod(d, m int) (div, mod int) {
	div = d / m
	mod = d % m
	return
}

func (b *Basic) PricePer15Minutes() int {
	return 220
}

func (b *Basic) MaxPrice() int {
	return 4290
}

func (m *Middle) PricePer15Minutes() int {
	return 330
}

func (m *Middle) MaxPrice() int {
	return 6490
}

func (p *Premium) PricePer15Minutes() int {
	return 440
}

func (p *Premium) MaxPrice() int {
	return 8690
}

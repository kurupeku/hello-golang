package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

type Middle struct{}

type Premium struct{}

func (basic Basic) PricePer15Minutes() int {
	return 220
}
func (middle Middle) PricePer15Minutes() int {
	return 330
}
func (premium Premium) PricePer15Minutes() int {
	return 440
}

func (basic Basic) MaxPrice() int {
	return 4290
}
func (middle Middle) MaxPrice() int {
	return 6490
}
func (premium Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	// TODO: 実装
	max_price_times := minutes / 360
	remain_minutes := minutes % 360
	price := car.MaxPrice() * max_price_times
	price_times := remain_minutes/15 + 1
	remain_price := car.PricePer15Minutes() * price_times
	if remain_price > car.MaxPrice() {
		remain_price = car.MaxPrice()
	}
	price += remain_price
	return price
}

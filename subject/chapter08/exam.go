package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

type Middle struct{}

type Premium struct{}

func Calc(car Car, minutes int) int {
	// TODO: 実装

	// 最大
	max_price_time := minutes / 360
	calc_max_price := max_price_time * car.MaxPrice()

	// 余り
	amari_time := minutes % 360
	calc_amari_time := amari_time % 15

	var calc_amari_price int
	var sum_price int

	switch calc_amari_time {
	case 0:
		calc_amari_price = (amari_time / 15) * car.PricePer15Minutes()
	default:
		calc_amari_price = ((amari_time / 15) + 1) * car.PricePer15Minutes()
	}

	switch {
	case calc_amari_price >= car.MaxPrice():
		calc_amari_price = car.MaxPrice()
	default:
		break
	}

	sum_price = calc_max_price + calc_amari_price

	return sum_price
}

func (Basic) PricePer15Minutes() int {
	return 220
}
func (Middle) PricePer15Minutes() int {
	return 330
}
func (Premium) PricePer15Minutes() int {
	return 440
}
func (Basic) MaxPrice() int {
	return 4290
}
func (Middle) MaxPrice() int {
	return 6490
}
func (Premium) MaxPrice() int {
	return 8690
}

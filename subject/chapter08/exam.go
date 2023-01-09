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
	// 最大時間360
	// 最大料金
	// 料金=(minutes / 15)の整数部分 * PricePer15Minutes
	// if 料金 >= MaxPrice
	//
	max_price_times := minutes / 360
	price := car.MaxPrice() * max_price_times

	amari := minutes % 360
	price_times := amari/15 + 1
	amari_price := price_times * car.PricePer15Minutes()
	if amari_price > car.MaxPrice() {
		amari_price = car.MaxPrice()
	}
	price += amari_price
	return price

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

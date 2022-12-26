package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

func (basic Basic) PricePer15Minutes() int {
	return 220
}

func (basic Basic) MaxPrice() int {
	return 4290
}

type Middle struct{}

func (middle Middle) PricePer15Minutes() int {
	return 330
}

func (middle Middle) MaxPrice() int {
	return 6490
}

type Premium struct{}

func (premium Premium) PricePer15Minutes() int {
	return 440
}

func (premium Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {

	hours := minutes / 60
	if hours >= 6 {
		resetCount := minutes / 360
		overMinutes := minutes - resetCount*360
		pricePer15Minutes := overMinutes / 15
		basePrice := (pricePer15Minutes + 1) * car.PricePer15Minutes()
		if basePrice >= car.MaxPrice() {
			return car.MaxPrice() + resetCount*car.MaxPrice()
		}
		return resetCount*car.MaxPrice() + basePrice
	}

	pricePer15Minutes := minutes / 15
	basePrice := (pricePer15Minutes + 1) * car.PricePer15Minutes()
	if hours < 6 && basePrice > car.MaxPrice() {
		return car.MaxPrice()
	}
	if pricePer15Minutes+1 > 0 {
		return (pricePer15Minutes + 1) * car.PricePer15Minutes()
	}

	return basePrice
}

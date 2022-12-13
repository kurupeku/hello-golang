package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

func (basic Basic) PricePer15Minutes() int {
	return 220;
}

func (basic Basic) MaxPrice() int {
	return 4290;
}

type Middle struct{}

func (middle Middle) PricePer15Minutes() int {
	return 330;
}

func (middle Middle) MaxPrice() int {
	return 6490;
}

type Premium struct{}

func (premium Premium) PricePer15Minutes() int {
	return 440;
}

func (premium Premium) MaxPrice() int {
	return 8690;
}

func getTotalPricePer15(pricePer15Minutes int, minutes int) int {
	return (minutes / 15 + 1) * pricePer15Minutes
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Calc(car Car, minutes int) int {
	hour6 := 360
	countOver6hour := minutes/hour6
	pricePerMaxPrice := car.MaxPrice() * countOver6hour
	remainingPrice := getTotalPricePer15(car.PricePer15Minutes(), (minutes - hour6 * countOver6hour))
	price := pricePerMaxPrice + min(remainingPrice, car.MaxPrice())
	return price
}

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
	return 0
}

package chapter08

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

// Car interface を実装するようにメソッドを定義し、その実装してください
type Basic struct{}

// Car interface を実装するようにメソッドを定義し、その実装してください
type Middle struct{}

// Car interface を実装するようにメソッドを定義し、その実装してください
type Premium struct{}

// 渡されたクルマのグレードと駐車時間から利用料金を計算する関数を実装してください
func Calc(car Car, minutes int) int {
	// TODO: 実装
	return 0
}

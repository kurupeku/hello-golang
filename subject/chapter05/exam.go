package chapter05

import (
	"fmt"
	"math"
)

func DarumaDrop(daruma []int) []int {
	// TODO: 実装

	//darumaの要素を取得し、長さとしてlengthとして定義する
	var length = len(daruma)
	//darumaの長さが1以下の場合
	if length <= 1 {
		//そのままdarumaを返す
		return daruma
	}

	//除外するdaruma(eliminated_daruma)はdarumaの長さ(length)を2分の1にしたものなる
	//長さ(length)は小数点で出てくるので、関数math.Ceilを使って小数点を繰り上げる
	//この時点では、eliminated_darumaは小数(float64)扱いとなるので要注意
	eliminated_daruma := math.Ceil(float64(length) / 2)
	//最終的なdaruma(final_daruma)としてint(整数)で定義し、indexを作る
	//final_darumaは、eliminated_darumaから-1（真ん中をすっ飛ばした）したものとなる
	//eliminated_darumaは小数(float64)扱いとなるので、整数(int)へ変換してあげる
	final_daruma := int(eliminated_daruma) - 1

	//真ん中のdarumaをすっ飛ばして残ったdarumaを格納する箱(empty_daruma)を作る
	//最終的なdarumaを格納する箱が必要となるので、新しい空のスライスを作る
	//新しいスライスはint(整数)を格納するのでintで問題なし
	var empty_daruma []int

	//iにループ回数をvに配列darumaの値を代入する
	//範囲はexam_test.goで定義されているdarumaの範囲内となる
	for i, v := range daruma {
		//index(final_daruma)にならない限りi回分、ループする
		if i != final_daruma {
			//空の箱(empty_daruma)にdarumaの値を追加し続ける
			empty_daruma = append(empty_daruma, v)
		}
	}
	//最後に残ったdaruma(empty_darumaに格納されたdaruma)を返す
	return empty_daruma
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装

	//配列(Hairetsu)というスライスをmakeを使用して作成する
	//長さ(len)はseedから呼び出す
	Hairetsu := make([][]int, len(seed))
	//長さ(len(seed))が0の時には、
	if len(seed) == 0 {
		//そのままスライスを返す
		return [][]int{}
	}
	//iは配列番号、vは掛け合わせる範囲を指定する
	//範囲はseedから呼び出して、for文で回す
	for i, v := range seed {
		//配列(Hairetsu)というスライスをmakeを使用して作成する
		//長さはlen(seed)にて定義、3個の配列を呼び出す
		Hairetsu[i] = make([]int, len(seed))
		//jは配列番号、xは掛け合わせる範囲を指定する
		//範囲はseedから呼び出して、for文で回す
		for j, x := range seed {
			//内側の配列同士を掛け合わせる
			//1*1,1*2,1*3, 2*1,2*2,2*3, 3*1,3*2,3*3
			//iとjは配列の番号
			Hairetsu[i][j] = v * x
		}
	}
	//配列（Hairetsu）を表示する
	fmt.Println(Hairetsu)
	//配列（Hairetsu）を返す
	return Hairetsu
}

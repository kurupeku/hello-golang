package chapter03

import "github.com/kurupeku/hello-golang/helper"

// If を使用して料金の条件分岐を行ってください
// 内回りと外回りで最安をだす問題と勘違いした
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	kyori := 0
	kane := 0
	kyoriinner := 0
	//kyoriouter := 0
	imaekiinner := "東京"
	//imaekiouter := "東京"

	if station == "東京" {
		kane = 0
		return kane
	}

	for {
		imaekiinner2 := helper.InnerNextStation(imaekiinner)
		//imaekiouter2 := helper.OuterNextStation(imaekiouter)
		kyoriinner += helper.InnerLoopDistance(imaekiinner2)
		//kyoriouter += helper.OuterLoopDistance(imaekiouter2)
		imaekiinner = imaekiinner2
		//imaekiouter = imaekiouter2
		if imaekiinner2 == station {
			kyori = kyoriinner
			break
		} // else if imaekiouter2 == station {
		//	kyori = kyoriouter
		//	break
		//}
	}

	if kyori >= 0 && kyori < 4000 {
		kane = 140
	} else if kyori >= 4000 && kyori < 7000 {
		kane = 160
	} else if kyori >= 7000 && kyori < 11000 {
		kane = 170
	} else if kyori >= 11000 && kyori < 16000 {
		kane = 200
	} else if kyori >= 16000 && kyori < 21000 {
		kane = 270
	} else if kyori >= 21000 && kyori < 26000 {
		kane = 350
	} else if kyori >= 26000 && kyori < 31000 {
		kane = 420
	} else {
		kane = 490
	}
	return kane
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	kyori := 0
	kane := 0
	//kyoriinner := 0
	kyoriouter := 0
	//imaekiinner := "東京"
	imaekiouter := "東京"
	if station == "東京" {
		kane = 0
		return kane
	}
	for {
		//imaekiinner2 := helper.InnerNextStation(imaekiinner)
		imaekiouter2 := helper.OuterNextStation(imaekiouter)
		//kyoriinner += helper.InnerLoopDistance(imaekiinner2)
		kyoriouter += helper.OuterLoopDistance(imaekiouter2)
		//imaekiinner = imaekiinner2
		imaekiouter = imaekiouter2
		//if imaekiinner2 == station {
		//	kyori = kyoriinner
		//	break
		//} else
		if imaekiouter2 == station {
			kyori = kyoriouter
			break
		}
	}

	switch {
	case kyori >= 0 && kyori < 4000:
		kane = 140
	case kyori >= 4000 && kyori < 7000:
		kane = 160
	case kyori >= 7000 && kyori < 11000:
		kane = 170
	case kyori >= 11000 && kyori < 16000:
		kane = 200
	case kyori >= 16000 && kyori < 21000:
		kane = 270
	case kyori >= 21000 && kyori < 26000:
		kane = 350
	case kyori >= 26000 && kyori < 31000:
		kane = 420
	case kyori >= 31000:
		kane = 490
	}
	return kane
}

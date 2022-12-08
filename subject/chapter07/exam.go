package chapter07

import (
	"fmt"
	"sort"
)

type Item struct {
	Name  string
	Price int
}

type Casher struct {
	List       map[string]int
	TotalPrice int
}

// c.List の Key を順番固定で返すメソッド
func (c *Casher) itemNames() []string {
	keys := make([]string, 0)
	for k := range c.List {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func (c *Casher) Purchase(item *Item) {
	// TODO: 実装
	// c.Listは最初nilでkeyが追加できないため、nilの場合はmakeで初期化するようにする
	if c.List == nil {
		c.List = make(map[string]int)
	}

	// 受け取ったitem.Nameをkeyとし、そのitem.Priceを追加（keyがある場合はPriceに加算される）
	c.List[item.Name] += item.Price

	// 受け取ったitem.Priceをc.TotalPriceに加算していく
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	// TODO: 実装
	// レシートの内容を記載していく変数を定義
	//var receipt string

	// レシートの初期化と店名の格納
	receipt := "\nラーメン道 楽酢\n\n"

	// c.Listに格納された商品名分ループ
	for _, itemName := range c.itemNames() {
		// 商品名(10桁の左詰め) + :(11桁目) + 商品に対応する合計金額(9桁の右詰め)
		receipt += fmt.Sprintf("%-10s:%9d\n", itemName, c.List[itemName])
	}

	// セパレート(20文字)
	receipt += "--------------------\n"

	// 合計金額(20桁の右詰め)
	receipt += fmt.Sprintf("%20d\n", c.TotalPrice)

	// レシートの内容を返す
	return receipt
}

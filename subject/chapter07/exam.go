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

	// c.Listがnilの時は初期化
	if c.List == nil {
		c.List = map[string]int{}
	}

	// すでに商品が存在している場合は、加算処理
	if _, ok := c.List[item.Name]; ok {
		c.List[item.Name] += item.Price
		c.TotalPrice += item.Price
		return
	}
	// 存在しない場合はマップに追加する
	c.List[item.Name] = item.Price
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	// TODO: 実装
	// c.Listをもとにレシートを生成する
	itemsNames := c.itemNames()
	s := fmt.Sprint("\n")
	s += fmt.Sprint("ラーメン道 楽酢\n")
	s += fmt.Sprintf("\n")
	for _, v := range itemsNames {
		s += fmt.Sprintf("%-10s:%9d\n", v, c.List[v])
	}
	s += fmt.Sprintf("--------------------\n")
	s += fmt.Sprintf("%20d", c.TotalPrice)
	s += fmt.Sprint("\n")
	fmt.Printf("%s", s)
	return s
}

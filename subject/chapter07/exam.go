package chapter07

import (
	"fmt"
	"sort"
	"strings"
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
	if c.List == nil {
		c.List = make(map[string]int, 3)
	}
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	// TODO: 実装
	receipt := "\n"
	receipt += "ラーメン道 楽酢\n"
	receipt += "\n"
	//ないよう繰り返し
	for _, v := range c.itemNames() {
		receipt += fmt.Sprintf("%-10s:%9d\n", v, c.List[v])
	}
	receipt += strings.Repeat("-", 20)
	receipt += "\n"
	receipt += fmt.Sprintf("%20d\n", c.TotalPrice)

	return receipt

}

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

	if c.List == nil {
		c.List = map[string]int{}
	}
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	// TODO: 実装

	names := c.itemNames()

	var receipt string

	receipt += "\nラーメン道 楽酢\n\n"

	for _, name := range names {
		receipt += fmt.Sprintf("%-10s:%9d\n", name, c.List[name])
	}

	receipt += "--------------------\n"
	receipt += fmt.Sprintf("%20d\n", c.TotalPrice)

	return receipt
}

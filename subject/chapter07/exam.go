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
	c.TotalPrice += item.Price
	if c.List == nil {
		c.List = make(map[string]int)
	}
	c.List[item.Name] += item.Price
}

func (c *Casher) Receipt() string {
	// TODO: 実装
	m := "\nラーメン道 楽酢\n\n"
	//m += "Ramen     :      890\n"
	for _, x := range c.itemNames() {
		m += fmt.Sprintf("%-10s:%9d\n", x, c.List[x])
	}
	m += "--------------------\n"
	m += fmt.Sprintf("%20d\n", c.TotalPrice)
	println(m)
	return m
}

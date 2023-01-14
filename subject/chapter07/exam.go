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

	switch {
	case c.List == nil:
		c.List = make(map[string]int)
	}

	switch {
	case c.List[item.Name] != 0:
		c.List[item.Name] = c.List[item.Name] + item.Price
	default:
		c.List[item.Name] = item.Price
	}

	total_price := &c.TotalPrice
	c.TotalPrice = *total_price + item.Price
}

func (c *Casher) Receipt() string {
	// TODO: 実装

	printer := "\nラーメン道 楽酢\n"
	for _, i := range c.itemNames() {
		printer = fmt.Sprintf(printer+"\n%-10s:%9d", i, c.List[i])
	}
	printer += "\n--------------------"
	printer = fmt.Sprintf(printer+"\n%20d\n", c.TotalPrice)

	return printer
}

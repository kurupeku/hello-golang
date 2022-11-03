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
	if c.List == nil {
		c.List = make(map[string]int)
	}
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	items := c.itemNames()
	str := "\nラーメン道 楽酢\n\n"
	for _, item := range items {
		str += fmt.Sprintf("%-10s%s%9d\n", item, ":", c.List[item])
	}
	str += "--------------------\n"
	str += fmt.Sprintf("%20d\n", c.TotalPrice)
	return str
}

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
		c.List = make(map[string]int, 0)
	}
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
	fmt.Println("item.Name", item.Name, "item total price", c.List[item.Name])
}

func (c *Casher) Receipt() string {
	r := "\nラーメン道 楽酢\n\n"

	for _, key := range c.itemNames() {
		r += fmt.Sprintf("%-10s:%9d\n", key, c.List[key])
	}

	r += "--------------------\n"
	r += fmt.Sprintf("%20d\n", c.TotalPrice)

	//fmt.Println(r)

	return r
}

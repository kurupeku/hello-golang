package chapter07

import (
	"fmt"
	"sort"
	"strconv"
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
	arr := make(map[string]int, 0)
	for k, v := range c.List {
		arr[k] = v
	}
	c.List = make(map[string]int, 0)
	for k, v := range arr {
		c.List[k] = v
	}

	if k, ok := c.List[item.Name]; !ok {
		c.List[item.Name] = item.Price
	} else {
		c.List[item.Name] = k + item.Price
	}
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	keys := c.itemNames()
	r := "\nラーメン道 楽酢\n"
	r += "\n"

	for _, v := range keys {
		r += fmt.Sprintf("%-10s:%9s\n", v, strconv.Itoa(c.List[v]))
	}

	r += "--------------------\n"
	r += fmt.Sprintf("%20s\n", strconv.Itoa(c.TotalPrice))
	return r
}

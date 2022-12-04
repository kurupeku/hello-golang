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
	if c.List == nil {
		c.List = make(map[string]int, 0)
	}
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	var sb strings.Builder
	sb.WriteString("\nラーメン道 楽酢\n\n")
	for _, val := range c.itemNames() {
		sb.WriteString(fmt.Sprintf("%-10s:%9d\n", val, c.List[val]))
	}
	sb.WriteString("--------------------\n")
	sb.WriteString(fmt.Sprintf("%20d\n", c.TotalPrice))
	return sb.String()
}

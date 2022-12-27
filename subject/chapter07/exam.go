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
		c.List = make(map[string]int)
	}
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {
	store_name := "\nラーメン道 楽酢\n"
	separator := "\n" + strings.Repeat("-", 20) + "\n"
	pre_result := make([]string, 0)
	for _, b := range c.itemNames() {
		tmp_result := "\n" + fmt.Sprintf("%-10s", b) + ":" + fmt.Sprintf("%9d", c.List[b])
		pre_result = append(pre_result, tmp_result)
	}
	var result string
	for _, d := range pre_result {
		result += d
	}

	return store_name + result + separator + fmt.Sprintf("%20d", c.TotalPrice) + "\n"
}

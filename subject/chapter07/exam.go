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
	// 明示的に初期化しないと怒られる
	if c.List == nil {
		c.List = make(map[string]int)
	}
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price

}

func (c *Casher) Receipt() string {
	// TODO: 実装
	header := "\nラーメン道 楽酢\n"
	//body := make([]string, 0)
	body := ""
	for _, order := range c.itemNames() {
		total := "\n" + fmt.Sprintf("%-10s", order) + ":" + fmt.Sprintf("%9d", c.List[order])
		body += total
	}
	//separate := fmt.Sprintf("My name is %s.", name)
	separate := "\n" + strings.Repeat("-", 20) + "\n"

	return header + body + separate + fmt.Sprintf("%20d", c.TotalPrice) + "\n"
}

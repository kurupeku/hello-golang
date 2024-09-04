package chapter07

import (
	"fmt"
	"sort"
	"strconv"
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

// 購入した商品 item の名前と金額を記録してください
func (c *Casher) Purchase(item *Item) {
	if c.List == nil {
		c.List = make(map[string]int)
	}

	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
}

// 購入した商品のレシートを規定のフォーマットに沿った文字列として返してください
func (c *Casher) Receipt() string {
	details := "\nラーメン道 楽酢\n\n"
	for _, n := range c.itemNames() {
		p := c.List[n]
		price := strconv.Itoa(p)
		details += fmt.Sprintf("%-10s: %8s\n", n, price)
	}

	sep := strings.Repeat("-", 20)
	tp := strconv.Itoa(c.TotalPrice)
	return details + fmt.Sprintf("%s\n%20s\n", sep, tp)
}

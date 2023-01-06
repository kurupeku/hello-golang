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

const Newline = "\n"
const Shopname = "ラーメン道 楽酢"
const Line = "-"

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

	c.Receipt()
}

func (c *Casher) Receipt() string {
	receipt := Newline
	receipt += Shopname
	receipt += strings.Repeat(Newline, 2)

	for _, value := range c.itemNames() {
		receipt += fmt.Sprintf("%-10s", value)
		receipt += ":"
		receipt += fmt.Sprintf("%9s", strconv.Itoa(c.List[value]))
		receipt += Newline
	}

	receipt += strings.Repeat(Line, 20)
	receipt += Newline
	receipt += fmt.Sprintf("%20s", strconv.Itoa(c.TotalPrice))
	receipt += Newline

	return receipt
}

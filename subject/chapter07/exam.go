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

	// 初期化されていなければ初期化
	if c.List == nil {
		c.List = map[string]int{}
	}

	// 受け取ったアイテムを加算
	c.List[item.Name] += item.Price
	c.TotalPrice += item.Price
}

func (c *Casher) Receipt() string {

	// TODO: 実装
	// ヘッダ部分
	receipt_str := "\nラーメン道 楽酢\n\n"

	// 品目
	for _, key := range c.itemNames() {
		receipt_str += fmt.Sprintf("%-10s:%9d\n", key, c.List[key])
	}
	// 合計値
	receipt_str += strings.Repeat("-", 20) + "\n" +
		fmt.Sprintf("%20d\n", c.TotalPrice)

	return receipt_str
}

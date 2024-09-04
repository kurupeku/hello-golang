package chapter07

import (
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

// 購入した商品 item の名前と金額を記録してください
func (c *Casher) Purchase(item *Item) {
	// TODO: 実装
}

// 購入した商品のレシートを規定のフォーマットに沿った文字列として返してください
func (c *Casher) Receipt() string {
	// TODO: 実装
	return ""
}

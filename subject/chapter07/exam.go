package chapter07

import (
	"sort"
	"fmt"
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

func (c *Casher) Purchase(item *Item) {
	if c.List == nil {
		c.List = make(map[string]int)
	}

	_, exist := c.List[item.Name];
	if exist {
		c.List[item.Name] += item.Price;
	} else {
		c.List[item.Name] = item.Price;
	}

	c.TotalPrice += item.Price;
}

func getTitle() string {
	title := "\n"
	title += fmt.Sprintf("%s\n", "ラーメン道 楽酢") 
	title += "\n"
	return title;
}

func getItemNameAndPrice(c *Casher) string {
	var itemNameAndPrice string

	MaxItemNameLength := 10
	MaxPriceLength := 9
	
	for _, itemName := range c.itemNames() {
		// 商品名を追加
		itemNameAndPrice += itemName

		// 商品名の後ろに半角スペースを追加
		itemNameAndPrice += strings.Repeat(" ", MaxItemNameLength - len(itemName))

		// :による区切り
		itemNameAndPrice += ":"

		// 商品価格の前に半角スペースを追加
		ItemPriceLength := len(strconv.Itoa(c.List[itemName]))
		itemNameAndPrice += strings.Repeat(" ", MaxPriceLength - ItemPriceLength)

		// 商品価格を追加
		itemNameAndPrice += strconv.Itoa(c.List[itemName])

		itemNameAndPrice += "\n"
	}

	return itemNameAndPrice;
}

func getTotalPrice(c *Casher) string {
	var totalPrice string;
	MaxLineLength := 20
	TotalPriceLength := len(strconv.Itoa(c.TotalPrice))
	totalPrice += strings.Repeat(" ", MaxLineLength - TotalPriceLength)
	totalPrice += strconv.Itoa(c.TotalPrice)
	totalPrice += "\n"
	return totalPrice;
}

func (c *Casher) Receipt() string {
	ans := getTitle()
	ans += getItemNameAndPrice(c)
	ans += "--------------------\n"
	ans += getTotalPrice(c)
	return ans
}

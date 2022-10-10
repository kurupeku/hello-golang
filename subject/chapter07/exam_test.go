package chapter07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCasher_Purchase(t *testing.T) {
	type args struct {
		item *Item
	}
	tests := []struct {
		name      string
		items     []Item
		wantList  map[string]int
		wantPrice int
	}{
		{
			name: "Purchase Ramen",
			items: []Item{
				{"Ramen", 890},
			},
			wantList: map[string]int{
				"Ramen": 890,
			},
			wantPrice: 890,
		},
		{
			name: "Purchase 2 Ramen",
			items: []Item{
				{"Ramen", 890},
				{"Ramen", 890},
			},
			wantList: map[string]int{
				"Ramen": 1780,
			},
			wantPrice: 1780,
		},
		{
			name: "Purchase 2 Ramen and Chahan",
			items: []Item{
				{"Ramen", 890},
				{"Ramen", 890},
				{"Chahan", 620},
			},
			wantList: map[string]int{
				"Ramen":  1780,
				"Chahan": 620,
			},
			wantPrice: 2400,
		},
		{
			name: "Purchase 2 Ramen, Chahan and 3 Gyoza",
			items: []Item{
				{"Ramen", 890},
				{"Ramen", 890},
				{"Chahan", 620},
				{"Gyoza", 520},
				{"Gyoza", 520},
				{"Gyoza", 520},
			},
			wantList: map[string]int{
				"Ramen":  1780,
				"Chahan": 620,
				"Gyoza":  1560,
			},
			wantPrice: 3960,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Casher{}
			for _, item := range tt.items {
				c.Purchase(&item)
			}
			assert.Equal(t, tt.wantList, c.List)
			assert.Equal(t, tt.wantPrice, c.TotalPrice)
		})
	}
}

func TestCasher_Receipt(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		want  string
	}{
		{
			name: "Purchase Ramen",
			items: []Item{},
			want: `
ラーメン道 楽酢

--------------------
                   0
`,
		},
		{
			name: "Purchase Ramen",
			items: []Item{
				{"Ramen", 890},
			},
			want: `
ラーメン道 楽酢

Ramen     :      890
--------------------
                 890
`,
		},
		{
			name: "Purchase 2 Ramen",
			items: []Item{
				{"Ramen", 890},
				{"Ramen", 890},
			},
			want: `
ラーメン道 楽酢

Ramen     :     1780
--------------------
                1780
`,
		},
		{
			name: "Purchase 2 Ramen and Chahan",
			items: []Item{
				{"Ramen", 890},
				{"Ramen", 890},
				{"Chahan", 620},
			},
			want: `
ラーメン道 楽酢

Chahan    :      620
Ramen     :     1780
--------------------
                2400
`,
		},
		{
			name: "Purchase 2 Ramen, Chahan and 3 Gyoza",
			items: []Item{
				{"Ramen", 890},
				{"Ramen", 890},
				{"Chahan", 620},
				{"Gyoza", 520},
				{"Gyoza", 520},
				{"Gyoza", 520},
			},
			want: `
ラーメン道 楽酢

Chahan    :      620
Gyoza     :     1560
Ramen     :     1780
--------------------
                3960
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Casher{}
			for _, item := range tt.items {
				c.Purchase(&item)
			}
			assert.Equal(t, tt.want, c.Receipt())
		})
	}
}

func TestCasher_itemNames(t *testing.T) {
	type fields struct {
		List map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Listed Ramen",
			fields: fields{map[string]int{
				"Ramen": 890,
			}},
			want: []string{"Ramen"},
		},
		{
			name: "Listed Ramen and Chahan",
			fields: fields{map[string]int{
				"Ramen":  890,
				"Chahan": 620,
			}},
			want: []string{"Chahan", "Ramen"},
		},
		{
			name: "Listed Ramen, Chahan and Gyoza",
			fields: fields{map[string]int{
				"Ramen":  890,
				"Chahan": 620,
				"Gyoza":  520,
			}},
			want: []string{"Chahan", "Gyoza", "Ramen"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Casher{
				List: tt.fields.List,
			}
			assert.Equal(t, tt.want, c.itemNames())
		})
	}
}

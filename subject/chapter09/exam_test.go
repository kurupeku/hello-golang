package chapter09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKaisatsu(t *testing.T) {
	type args struct {
		from    string
		to      string
		charger Charger
	}
	tests := []struct {
		name        string
		args        args
		want        bool
		wantBalance int
		wantPoint   int
		assertion   assert.ErrorAssertionFunc
	}{
		{
			name:        "新宿 to 新宿 using Card (Balance: 400, Point: 200)",
			args:        args{"新宿", "新宿", &Card{400, 200}},
			want:        false,
			wantBalance: 400,
			wantPoint:   200,
			assertion:   assert.Error,
		},
		{
			name:      "新宿 to 新宿 using Ticket (600)",
			args:      args{"新宿", "新宿", &Ticket{600, false}},
			want:      false,
			assertion: assert.Error,
		},
		{
			name:        "東京 to 新宿 using Card (Balance: 400, Point: 200)",
			args:        args{"東京", "新宿", &Card{400, 200}},
			want:        true,
			wantBalance: 336,
			wantPoint:   0,
			assertion:   assert.NoError,
		},
		{
			name:      "東京 to 新宿 using Ticket (600)",
			args:      args{"東京", "新宿", &Ticket{600, false}},
			want:      true,
			assertion: assert.NoError,
		},
		{
			name:        "東京 to 新宿 using Card (Balance: 100, Point: 0)",
			args:        args{"東京", "新宿", &Card{100, 0}},
			want:        false,
			wantBalance: 100,
			wantPoint:   0,
			assertion:   assert.NoError,
		},
		{
			name:      "東京 to 新宿 using Ticket (100)",
			args:      args{"東京", "新宿", &Ticket{100, false}},
			want:      false,
			assertion: assert.NoError,
		},
		{
			name:        "新宿 to 品川 using Card (Balance: 200, Point: 100)",
			args:        args{"新宿", "品川", &Card{200, 100}},
			want:        true,
			wantBalance: 132,
			wantPoint:   0,
			assertion:   assert.NoError,
		},
		{
			name:      "東京 to 新宿 using Ticket (600)",
			args:      args{"東京", "新宿", &Ticket{600, false}},
			want:      true,
			assertion: assert.NoError,
		},
		{
			name:        "新宿 to 品川 using Card (Balance: 100, Point: 50)",
			args:        args{"新宿", "品川", &Card{100, 50}},
			want:        false,
			wantBalance: 100,
			wantPoint:   50,
			assertion:   assert.NoError,
		},
		{
			name:      "東京 to 新宿 using Ticket (600)",
			args:      args{"新宿", "品川", &Ticket{120, false}},
			want:      false,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Kaisatsu(tt.args.from, tt.args.to, tt.args.charger)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			switch v := tt.args.charger.(type) {
			case *Card:
				assert.Equal(t, tt.wantBalance, v.Balance)
				assert.Equal(t, tt.wantPoint, v.Point)
			case *Ticket:
				assert.Equal(t, tt.want, v.Used)
			}
		})
	}
}

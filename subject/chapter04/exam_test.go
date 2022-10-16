package chapter04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKaisatsu(t *testing.T) {
	type args struct {
		charge int
		card   *Card
	}
	tests := []struct {
		name        string
		args        args
		want        bool
		wantBalance int
		wantPoint   int
	}{
		{
			name:        "Charge: 500, Balance: 1000, Point: 0, Use balance only",
			args:        args{charge: 500, card: &Card{1000, 0}},
			want:        true,
			wantBalance: 500,
			wantPoint:   0,
		},
		{
			name:        "Charge: 100, Balance: 50, Point: 200, Use point only",
			args:        args{charge: 100, card: &Card{50, 200}},
			want:        true,
			wantBalance: 50,
			wantPoint:   100,
		},
		{
			name:        "Charge: 800, Balance: 500, Point: 400, Use balance and point",
			args:        args{charge: 800, card: &Card{500, 400}},
			want:        true,
			wantBalance: 100,
			wantPoint:   0,
		},
		{
			name:        "Charge: 400, Balance: 200, Point: 200, Use all",
			args:        args{charge: 400, card: &Card{200, 200}},
			want:        true,
			wantBalance: 0,
			wantPoint:   0,
		},
		{
			name:        "Charge: 1000, Balance: 200, Point: 300, Could not pass Kaisatsu",
			args:        args{charge: 1000, card: &Card{200, 300}},
			want:        false,
			wantBalance: 200,
			wantPoint:   300,
		},
		{
			name:        "Charge: 300, Balance: 200, Point: 0, Could not pass Kaisatsu",
			args:        args{charge: 300, card: &Card{200, 0}},
			want:        false,
			wantBalance: 200,
			wantPoint:   0,
		},
		{
			name:        "Charge: 1000, Balance: 0, Point: 800, Could not pass Kaisatsu",
			args:        args{charge: 1000, card: &Card{0, 800}},
			want:        false,
			wantBalance: 0,
			wantPoint:   800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Kaisatsu(tt.args.charge, tt.args.card))
			assert.Equal(t, &Card{tt.wantBalance, tt.wantPoint}, tt.args.card)
		})
	}
}

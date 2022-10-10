package chapter03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInnerChargeFromTokyo(t *testing.T) {
	type args struct {
		station string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "To 東京",
			args: args{"東京"},
			want: 0,
		},
		{
			name: "To 神田",
			args: args{"神田"},
			want: 140,
		},
		{
			name: "To 日暮里",
			args: args{"日暮里"},
			want: 160,
		},
		{
			name: "To 巣鴨",
			args: args{"巣鴨"},
			want: 170,
		},
		{
			name: "To 池袋",
			args: args{"池袋"},
			want: 200,
		},
		{
			name: "To 新宿",
			args: args{"新宿"},
			want: 270,
		},
		{
			name: "To 大崎",
			args: args{"大崎"},
			want: 350,
		},
		{
			name: "To 品川",
			args: args{"品川"},
			want: 420,
		},
		{
			name: "To 新橋",
			args: args{"新橋"},
			want: 490,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, InnerChargeFromTokyo(tt.args.station))
		})
	}
}

func TestOuterChargeFromTokyo(t *testing.T) {
	type args struct {
		station string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "To 東京",
			args: args{"東京"},
			want: 0,
		},
		{
			name: "To 有楽町",
			args: args{"有楽町"},
			want: 140,
		},
		{
			name: "To 品川",
			args: args{"品川"},
			want: 160,
		},
		{
			name: "To 目黒",
			args: args{"目黒"},
			want: 170,
		},
		{
			name: "To 恵比寿",
			args: args{"恵比寿"},
			want: 200,
		},
		{
			name: "To 新宿",
			args: args{"新宿"},
			want: 270,
		},
		{
			name: "To 池袋",
			args: args{"池袋"},
			want: 350,
		},
		{
			name: "To 鶯谷",
			args: args{"鶯谷"},
			want: 420,
		},
		{
			name: "To 秋葉原",
			args: args{"秋葉原"},
			want: 490,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, OuterChargeFromTokyo(tt.args.station))
		})
	}
}

package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInnerNextStation(t *testing.T) {
	type args struct {
		current string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Current in 秋葉原",
			args: args{"秋葉原"},
			want: "御徒町",
		},
		{
			name: "Current in 新宿",
			args: args{"新宿"},
			want: "代々木",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, InnerNextStation(tt.args.current))
		})
	}
}

func TestInnerLoopDistance(t *testing.T) {
	type args struct {
		station string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Throw a proper argument",
			args: args{"神田"},
			want: 1300,
		},
		{
			name: "Throw a improper argument",
			args: args{"大阪"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, InnerLoopDistance(tt.args.station))
		})
	}
}

func TestOuterNextStation(t *testing.T) {
	type args struct {
		current string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Current in 品川",
			args: args{"品川"},
			want: "大崎",
		},
		{
			name: "Current in 恵比寿",
			args: args{"恵比寿"},
			want: "渋谷",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, OuterNextStation(tt.args.current))
		})
	}
}

func TestOuterLoopDistance(t *testing.T) {
	type args struct {
		station string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Throw a proper argument",
			args: args{"神田"},
			want: 700,
		},
		{
			name: "Throw a improper argument",
			args: args{"大阪"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, OuterLoopDistance(tt.args.station))
		})
	}
}

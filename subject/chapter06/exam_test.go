package chapter06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAhoNumber_Call(t *testing.T) {
	tests := []struct {
		name string
		n    AhoNumber
		want string
	}{
		{
			name: "throw 1",
			n:    1,
			want: "1",
		},
		{
			name: "throw 3",
			n:    3,
			want: "3 aho",
		},
		{
			name: "throw 13",
			n:    13,
			want: "13 aho",
		},
		{
			name: "throw 19",
			n:    19,
			want: "19",
		},
		{
			name: "throw 131",
			n:    131,
			want: "131 aho",
		},
		{
			name: "throw 198",
			n:    198,
			want: "198 aho",
		},
		{
			name: "throw 202",
			n:    202,
			want: "202",
		},
		{
			name: "throw 1038",
			n:    1038,
			want: "1038 aho",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.n.Call())
		})
	}
}

func TestNabeatsu(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "throw 3",
			args: args{3},
			want: []string{"1", "2", "3 aho"},
		},
		{
			name: "throw 10",
			args: args{10},
			want: []string{"1", "2", "3 aho", "4", "5", "6 aho", "7", "8", "9 aho", "10"},
		},
		{
			name: "throw 40",
			args: args{40},
			want: []string{
				"1", "2", "3 aho", "4", "5", "6 aho", "7", "8", "9 aho", "10",
				"11", "12 aho", "13 aho", "14", "15 aho", "16", "17", "18 aho", "19", "20",
				"21 aho", "22", "23 aho", "24 aho", "25", "26", "27 aho", "28", "29", "30 aho",
				"31 aho", "32 aho", "33 aho", "34 aho", "35 aho", "36 aho", "37 aho", "38 aho", "39 aho", "40",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Nabeatsu(tt.args.n))
		})
	}
}

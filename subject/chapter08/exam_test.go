package chapter08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	type args struct {
		car     Car
		minutes int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Use Basic for 40 minutes",
			args: args{&Basic{}, 40},
			want: 660,
		},
		{
			name: "Use Middle for 118 minutes",
			args: args{&Middle{}, 118},
			want: 2640,
		},
		{
			name: "Use Premium for 252 minutes",
			args: args{&Premium{}, 252},
			want: 7480,
		},
		{
			name: "Use Basic for 335 minutes and reach max price",
			args: args{&Basic{}, 335},
			want: 4290,
		},
		{
			name: "Use Middle for 933 minutes and 2 times reaching max price",
			args: args{&Middle{}, 933},
			want: 17930,
		},
		{
			name: "Use Premium for 2142 minutes and 5 times reaching max price",
			args: args{&Premium{}, 2142},
			want: 52140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Calc(tt.args.car, tt.args.minutes))
		})
	}
}

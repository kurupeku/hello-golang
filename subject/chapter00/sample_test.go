package chapter00

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 + 1",
			args: args{a: 1, b: 1},
			want: 2,
		},
		{
			name: "5 - 7",
			args: args{a: 5, b: -7},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Add(tt.args.a, tt.args.b))
		})
	}
}

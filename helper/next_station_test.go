package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			name: "throw a proper argument",
			args: args{"神田"},
			want: 1300,
		},
		{
			name: "throw a improper argument",
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
			name: "throw a proper argument",
			args: args{"神田"},
			want: 700,
		},
		{
			name: "throw a improper argument",
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

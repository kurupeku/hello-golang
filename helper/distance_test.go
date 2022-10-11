package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDistance(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "From km",
			args: args{"1.55km"},
			want: 1550,
		},
		{
			name: "From m",
			args: args{"2500m"},
			want: 2500,
		},
		{
			name: "From cm",
			args: args{"4800cm"},
			want: 48,
		},
		{
			name: "From mm",
			args: args{"790000mm"},
			want: 790,
		},
		{
			name: "From integer",
			args: args{"4739"},
			want: 4739,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ParseDistance(tt.args.d))
		})
	}
}

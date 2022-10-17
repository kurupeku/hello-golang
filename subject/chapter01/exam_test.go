package chapter01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxi(t *testing.T) {
	type args struct {
		distance string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "Distance is 1700",
			args:  args{"1700"},
			want:  500,
			want1: 600,
		},
		{
			name:  "Distance is 34500",
			args:  args{"34500"},
			want:  13700,
			want1: 16440,
		},
		{
			name:  "Distance is 8900m",
			args:  args{"8900m"},
			want:  3400,
			want1: 4080,
		},
		{
			name:  "Distance is 1500m",
			args:  args{"1500m"},
			want:  500,
			want1: 600,
		},
		{
			name:  "Distance is 8900m",
			args:  args{"8900m"},
			want:  3400,
			want1: 4080,
		},
		{
			name:  "Distance is 1.8km",
			args:  args{"1.8km"},
			want:  600,
			want1: 720,
		},
		{
			name:  "Distance is 56km",
			args:  args{"56km"},
			want:  22300,
			want1: 26760,
		},
		{
			name:  "Distance is 360000cm",
			args:  args{"360000cm"},
			want:  1300,
			want1: 1560,
		},
		{
			name:  "Distance is 9240000cm",
			args:  args{"9240000cm"},
			want:  36800,
			want1: 44160,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Taxi(tt.args.distance)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

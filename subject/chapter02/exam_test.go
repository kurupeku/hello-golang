package chapter02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumCoins(t *testing.T) {
	type args struct {
		price uint
	}
	tests := []struct {
		name         string
		args         args
		wantCount500 uint
		wantCount100 uint
		wantCount050 uint
		wantCount010 uint
		wantCount005 uint
		wantCount001 uint
	}{
		{
			name:         "Throw 298",
			args:         args{298},
			wantCount500: 0,
			wantCount100: 3,
			wantCount050: 0,
			wantCount010: 2,
			wantCount005: 1,
			wantCount001: 2,
		},
		{
			name:         "Throw 1907",
			args:         args{1907},
			wantCount500: 4,
			wantCount100: 0,
			wantCount050: 1,
			wantCount010: 4,
			wantCount005: 1,
			wantCount001: 2,
		},
		{
			name:         "Throw 87",
			args:         args{87},
			wantCount500: 0,
			wantCount100: 0,
			wantCount050: 1,
			wantCount010: 4,
			wantCount005: 1,
			wantCount001: 0,
		},
		{
			name:         "Throw 15801",
			args:         args{15801},
			wantCount500: 34,
			wantCount100: 3,
			wantCount050: 1,
			wantCount010: 3,
			wantCount005: 0,
			wantCount001: 1,
		},
		{
			name:         "Throw 12",
			args:         args{12},
			wantCount500: 0,
			wantCount100: 0,
			wantCount050: 0,
			wantCount010: 1,
			wantCount005: 0,
			wantCount001: 3,
		},
		{
			name:         "Throw 495",
			args:         args{495},
			wantCount500: 1,
			wantCount100: 0,
			wantCount050: 0,
			wantCount010: 4,
			wantCount005: 0,
			wantCount001: 4,
		},
		{
			name:         "Throw 0",
			args:         args{0},
			wantCount500: 0,
			wantCount100: 0,
			wantCount050: 0,
			wantCount010: 0,
			wantCount005: 0,
			wantCount001: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount500, gotCount100, gotCount050, gotCount010, gotCount005, gotCount001 := MinimumCoins(tt.args.price)
			assert.Equal(t, tt.wantCount500, gotCount500)
			assert.Equal(t, tt.wantCount100, gotCount100)
			assert.Equal(t, tt.wantCount050, gotCount050)
			assert.Equal(t, tt.wantCount010, gotCount010)
			assert.Equal(t, tt.wantCount005, gotCount005)
			assert.Equal(t, tt.wantCount001, gotCount001)
		})
	}
}

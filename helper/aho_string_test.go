package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAhoString(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "throw 3",
			args: args{3},
			want: "さぁん",
		},
		{
			name: "throw 10",
			args: args{10},
			want: "じゅう",
		},
		{
			name: "throw 13",
			args: args{13},
			want: "じゅうさぁん",
		},
		{
			name: "throw 35",
			args: args{35},
			want: "さんじゅうごぉ",
		},
		{
			name: "throw 91",
			args: args{91},
			want: "きゅうじゅういぃち",
		},
		{
			name: "throw 213",
			args: args{213},
			want: "にひゃくじゅうさぁん",
		},
		{
			name: "throw 882",
			args: args{882},
			want: "はっぴゃくはちじゅうにぃ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AhoString(tt.args.i))
		})
	}
}

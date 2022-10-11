package chapter05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDarumaDrop(t *testing.T) {
	type args struct {
		daruma []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Daruma is empty",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "Daruma is 1",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "Daruma is 1, 2",
			args: args{[]int{1, 2}},
			want: []int{2},
		},
		{
			name: "Daruma is 1, 2, 3",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "Daruma is 1, 2, 3, 4",
			args: args{[]int{1, 2, 3, 4}},
			want: []int{1, 3, 4},
		},
		{
			name: "Daruma is 1, 2, 3, 4, 5",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "Daruma is 1, 2, 3, 4, 5, 6",
			args: args{[]int{1, 2, 3, 4, 5, 6}},
			want: []int{1, 2, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DarumaDrop(tt.args.daruma))
		})
	}
}

func TestMatrixMultiple(t *testing.T) {
	type args struct {
		seed []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Throw empty slice",
			args: args{[]int{}},
			want: [][]int{},
		},
		{
			name: "Throw 1",
			args: args{[]int{1}},
			want: [][]int{
				{1},
			},
		},
		{
			name: "Throw 1, 2, 3",
			args: args{[]int{1, 2, 3}},
			want: [][]int{
				{1, 2, 3},
				{2, 4, 6},
				{3, 6, 9},
			},
		},
		{
			name: "Throw 5, -10, 15",
			args: args{[]int{5, -10, 15}},
			want: [][]int{
				{25, -50, 75},
				{-50, 100, -150},
				{75, -150, 225},
			},
		},
		{
			name: "Throw 23, 180, 403",
			args: args{[]int{23, 180, 403}},
			want: [][]int{
				{529, 4140, 9269},
				{4140, 32400, 72540},
				{9269, 72540, 162409},
			},
		},
		{
			name: "Throw -3, 8, -14, 85, -128",
			args: args{[]int{-3, 8, -14, 85, -128}},
			want: [][]int{
				{9, -24, 42, -255, 384},
				{-24, 64, -112, 680, -1024},
				{42, -112, 196, -1190, 1792},
				{-255, 680, -1190, 7225, -10880},
				{384, -1024, 1792, -10880, 16384},
			},
		},
		{
			name: "Throw 328, 23, 52",
			args: args{[]int{328, 23, 52}},
			want: [][]int{
				{107584, 7544, 17056},
				{7544, 529, 1196},
				{17056, 1196, 2704},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MatrixMultiple(tt.args.seed))
		})
	}
}

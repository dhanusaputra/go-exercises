package main

import (
	"reflect"
	"testing"
)

func TestGetColatz(t *testing.T) {
	type args struct {
		n int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "happy path",
			args: args{
				n: 4,
			},
			want: []int{4, 2, 1},
		},
		{
			name: "happy path 2",
			args: args{
				n: 40,
			},
			want: []int{40, 20, 10, 5, 16, 8, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getColatz(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColatz() = %v, want %v", got, tt.want)
			}
		})
	}
}

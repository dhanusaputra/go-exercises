package main

import (
	"reflect"
	"testing"
)

func TestGetTriangularNumber(t *testing.T) {
	type args struct {
		n int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path 1",
			args: args{
				n: 4,
			},
			want: 10,
		},
		{
			name: "happy path 2",
			args: args{
				n: 50,
			},
			want: 1275,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTringularNumber(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTriangularNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDivisors(t *testing.T) {
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
				n: 155,
			},
			want: []int{1, 5, 31, 155},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDivisors(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

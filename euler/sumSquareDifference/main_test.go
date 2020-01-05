package main

import (
	"reflect"
	"testing"
)

func TestSumOfSquare(t *testing.T) {
	type args struct {
		n int
	}
	var tests = []struct {
		name string
		args args
		want float64
	}{
		{
			name: "happy path",
			args: args{
				n: 5,
			},
			want: 55,
		},
		{
			name: "happy path 2",
			args: args{
				n: 50,
			},
			want: 42925,
		},
		{
			name: "happy path 500",
			args: args{
				n: 500,
			},
			want: 4.179175e+07,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumOfSquare(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquareOfSum(t *testing.T) {
	type args struct {
		n int
	}
	var tests = []struct {
		name string
		args args
		want float64
	}{
		{

			name: "happy path",
			args: args{
				n: 5,
			},
			want: 225,
		},
		{
			name: "happy path 2",
			args: args{
				n: 50,
			},
			want: 1.625625e+06,
		},
		{
			name: "happy path 500",
			args: args{
				n: 500,
			},
			want: 1.56875625e+10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := squareOfSum(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

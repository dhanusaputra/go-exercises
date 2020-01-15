package main

import (
	"reflect"
	"testing"
)

func TestGetDivisors(t *testing.T) {
	type args struct {
		n float64
	}
	var tests = []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "happy path",
			args: args{
				n: 220,
			},
			want: []float64{1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDivisors(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDivisors() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

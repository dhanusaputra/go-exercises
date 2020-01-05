package main

import (
	"reflect"
	"testing"
)

func TestPythagoreanTriplet(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path",
			args: args{
				a: 3,
				b: 4,
				c: 5,
			},
			want: true,
		},
		{
			name: "happy path not pythagorean triplet",
			args: args{
				a: 5,
				b: 4,
				c: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPythagoreanTriplet(tt.args.a, tt.args.b, tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsPythagoreanTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}

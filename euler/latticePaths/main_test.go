package main

import (
	"reflect"
	"testing"
)

func TestCountingLaticePaths(t *testing.T) {
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
				n: 10,
			},
			want: 184756,
		},
		{
			name: "happy path 2",
			args: args{
				n: 22,
			},
			want: 2.1040989637199998e+12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countingLaticePaths(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountingLaticePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

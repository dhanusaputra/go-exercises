package main

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path prime",
			args: args{
				n: 2,
			},
			want: true,
		},
		{
			name: "happy path not prime",
			args: args{
				n: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPrime(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

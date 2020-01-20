package main

import "testing"

import "reflect"

func TestJoin(t *testing.T) {
	type args struct {
		r []rune
		c rune
	}
	var tests = []struct {
		name string
		args args
		want []string
	}{
		{"happy path", args{[]rune("abcd"), 'x'}, []string{"xabcd", "axbcd", "abxcd", "abcxd", "abcdx"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := join(tt.args.r, tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}

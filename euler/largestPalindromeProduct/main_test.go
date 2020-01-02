package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		n int
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path palindrome",
			args: args{
				n: 434,
			},
			want: true,
		},
		{
			name: "happy path not palindrome",
			args: args{
				n: 432,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInt(t *testing.T) {
	type args struct {
		n int
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path 1",
			args: args{
				n: 234,
			},
			want: "432",
		},
		{
			name: "happy path 2",
			args: args{
				n: 124156,
			},
			want: "651421",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseInt(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

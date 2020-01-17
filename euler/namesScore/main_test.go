package main

import (
	"reflect"
	"testing"
)

func TestGetStringWorth(t *testing.T) {
	var tests = []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "happy path",
			arg:  "GOLANG",
			want: 56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getStringWorth(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStringWorth() = %v, want %v", got, tt.want)
			}
		})
	}
}

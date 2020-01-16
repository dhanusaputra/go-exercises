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

func TestIsAbundantNumber(t *testing.T) {
	var tests = []struct {
		name string
		arg  float64
		want bool
	}{
		{
			name: "happy path return true",
			arg:  12,
			want: true,
		},
		{
			name: "happy path return false",
			arg:  13,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAbundantNumber(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isAbundantNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSumBetweenSlices(t *testing.T) {
	type args struct {
		s1 []float64
		s2 []float64
	}
	var tests = []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "happy path",
			args: args{
				s1: []float64{1, 3, 5},
				s2: []float64{2, 4},
			},
			want: []float64{3, 5, 5, 7, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSumBetweenSlices(tt.args.s1, tt.args.s2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSumBetweenSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContains(t *testing.T) {
	type args struct {
		s []float64
		f float64
	}
	var tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path return true",
			args: args{
				s: []float64{1, 3, 5},
				f: 3,
			},
			want: true,
		},
		{
			name: "happy path return false",
			args: args{
				s: []float64{1, 3, 5},
				f: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceContains(tt.args.s, tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

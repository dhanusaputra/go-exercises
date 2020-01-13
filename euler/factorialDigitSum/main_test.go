package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetFactorial(t *testing.T) {
	type args struct {
		n int64
	}
	var tests = []struct {
		name       string
		args       args
		want       big.Int
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "happy path",
			args: args{
				n: 10,
			},
			want:       *big.NewInt(3628800),
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFactorial(tt.args.n)
			if tt.wantErr && err == nil {
				t.Errorf("getFactorial() err = %+v, wantErr %+v", err, tt.wantErr)
			}
			if tt.wantErr && !reflect.DeepEqual(err.Error(), tt.wantErrMsg) {
				t.Errorf("getFactorial() err = %+v, wantErrMsg %+v", err.Error(), tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFactorial() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
func TestSumOfCharDigits(t *testing.T) {
	type args struct {
		s string
	}
	var tests = []struct {
		name       string
		args       args
		want       int
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "happy path",
			args: args{
				s: "102394",
			},
			want:       19,
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sumOfCharDigits(tt.args.s)
			if tt.wantErr && err == nil {
				t.Errorf("sumOfCharDigits() err = %+v, wantErr %+v", err, tt.wantErr)
			}
			if tt.wantErr && !reflect.DeepEqual(err.Error(), tt.wantErrMsg) {
				t.Errorf("sumOfCharDigits() err = %+v, wantErrMsg %+v", err.Error(), tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfCharDigits() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestIsEvenlyDivisible(t *testing.T) {
	type args struct {
		n   int
		div int
	}
	var tests = []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
		want       bool
	}{
		{
			name: "happy path",
			args: args{
				n:   10,
				div: 2,
			},
			wantErr:    false,
			wantErrMsg: "",
			want:       true,
		},
		{
			name: "happy path not divisible",
			args: args{
				n:   5,
				div: 2,
			},
			wantErr:    false,
			wantErrMsg: "",
			want:       false,
		},
		{
			name: "unhappy path",
			args: args{
				n:   -10,
				div: 2,
			},
			wantErr:    true,
			wantErrMsg: "Error number is not positive",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isEvenlyDivisible(tt.args.n, tt.args.div)
			if err != nil && !tt.wantErr {
				t.Errorf("IsEvenlyDivisible() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.wantErrMsg {
				t.Errorf("IsEvenlyDivisible() error message = %s, wantErrMsg %s", err.Error(), tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEvenlyDivisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

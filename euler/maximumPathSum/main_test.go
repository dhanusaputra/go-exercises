package main

import (
	"reflect"
	"testing"
)

func TestParseStringToFloatSlice(t *testing.T) {
	type args struct {
		s string
	}
	var tests = []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
		want       [][]float64
	}{
		{
			name: "happy path",
			args: args{
				s: `75
				95 64
				17 47 82`,
			},
			wantErr:    false,
			wantErrMsg: "",
			want:       [][]float64{{75}, {95, 64}, {17, 47, 82}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStringToFloatSlice(tt.args.s)
			if tt.wantErr && err != nil {
				t.Errorf("parseStringToFloatSlice error = %+v, wantErr %+v", err, tt.wantErr)
			}
			if tt.wantErr && !reflect.DeepEqual(err.Error(), tt.wantErrMsg) {
				t.Errorf("parseStringToFloatSlice error = %+v, wantErrMsg %+v", err.Error(), tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStringToFloatSlice = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSumMaximumPath(t *testing.T) {
	type args struct {
		f [][]float64
	}
	var tests = []struct {
		name       string
		args       args
		want       float64
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "happy path",
			args: args{
				f: [][]float64{{75}, {95, 64}, {17, 47, 82}},
			},
			want:       221,
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sumMaximumPath(tt.args.f)
			if tt.wantErr && err != nil {
				t.Errorf("sumMaximumPath() err = %+v, wantErr %+v", err, tt.wantErr)
			}
			if tt.wantErr && !reflect.DeepEqual(err.Error(), tt.wantErrMsg) {
				t.Errorf("sumMaximumPath() err = %+v, wantErrMsg %+v", err, tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumMaximumPath() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

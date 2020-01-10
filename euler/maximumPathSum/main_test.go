package main

import (
	"reflect"
	"testing"
)

func TestParseStringToSlice(t *testing.T) {
	type args struct {
		s string
	}
	var tests = []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
		want       [][]string
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
			want:       [][]string{{"75"}, {"95", "64"}, {"17", "47", "82"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStringToSlice(tt.args.s)
			if tt.wantErr && err != nil {
				t.Errorf("ParseStringToSlice error = %+v, wantErr %+v", err, tt.wantErr)
			}
			if tt.wantErr && !reflect.DeepEqual(err.Error(), tt.wantErrMsg) {
				t.Errorf("ParseStringToSlice error = %+v, wantErrMsg %+v", err.Error(), tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringToSlice = %+v, want %+v", got, tt.want)
			}
		})
	}
}

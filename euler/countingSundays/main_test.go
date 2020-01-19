package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCountSundayInFirstMonth(t *testing.T) {
	type args struct {
		s time.Time
		e time.Time
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
				s: getTime("2000-01-01"),
				e: getTime("2020-01-01"),
			},
			want:       35,
			wantErr:    false,
			wantErrMsg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countSundayInFirstMonth(tt.args.s, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("countSundayInFirstMonth() err = %+v, wantErr %+v", err, tt.wantErr)
			}
			if tt.wantErr && !reflect.DeepEqual(err.Error(), tt.wantErrMsg) {
				t.Errorf("countSundayInFirstMonth() errMsg = %+v, wantErrMsg %+v", err.Error(), tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSundayInFirstMonth() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func getTime(s string) time.Time {
	res, err := time.Parse(time.RFC3339[0:10], s)
	if err != nil {
		panic(err)
	}
	return res
}

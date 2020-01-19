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

func TestParseTxt(t *testing.T) {
	var tests = []struct {
		name       string
		arg        string
		want       []string
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:       "happy path",
			arg:        "testdata/p022_names.txt",
			want:       []string{`"MARY","PATRICIA","LINDA","BARBARA","ELIZABETH","JENNIFER","MARIA","SUSAN","MARGARET","DOROTHY","LISA","NANCY","KAREN","BETTY"`},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name:       "unhappy path",
			arg:        "testdata/p023_names.txt",
			want:       nil,
			wantErr:    true,
			wantErrMsg: "open testdata/p023_names.txt: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTxt(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTxt() err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && !reflect.DeepEqual(err.Error(), tt.wantErrMsg) {
				t.Errorf("parseTxt() err = %v, wantErrMsg %v", err.Error(), tt.wantErrMsg)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTxt() = %v, want %v", got, tt.want)
			}
		})
	}
}

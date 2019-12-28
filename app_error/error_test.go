package app_error

import (
	"reflect"
	"strings"
	"testing"
)

func TestErrHandling(t *testing.T) {
	type args struct {
		typeS string
	}
	tests := []struct {
		name        string
		args        args
		wantErrType string
	}{
		{name: "HogeErr", args: args{typeS: "hoge"}, wantErrType: "HogeError"},
		{name: "FooErr", args: args{typeS: "foo"}, wantErrType: "FooError"},
		{name: "BarErr", args: args{typeS: "bar"}, wantErrType: "BarError"},
		{name: "UnknownErr1", args: args{typeS: "moge"}, wantErrType: "UnknownError"},
		{name: "UnknownErr2", args: args{typeS: "fuga"}, wantErrType: "UnknownError"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewAppErr(tt.args.typeS)
			split := strings.Split(reflect.ValueOf(err).Type().String(), ".")
			errType := split[len(split)-1]
			if errType != tt.wantErrType {
				t.Log("error type is unexpected", errType, tt.wantErrType)
				t.Fail()
			}
		})
	}
}

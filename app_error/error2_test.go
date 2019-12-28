package app_error

import "testing"

func TestIsBadRequest(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ValidationErr", args: args{err: NewAppErr2("ValidationErr")}, want: true},
		{name: "WrappedValidationErr", args: args{err: NewAppErr2("WrappedValidationErr")}, want: true},
		{name: "DuplicatedIdErr", args: args{err: NewAppErr2("DuplicatedIdErr")}, want: true},
		{name: "UnauthorizedErr", args: args{err: NewAppErr2("UnauthorizedErr")}, want: false},
		{name: "ResourceNotFoundErr", args: args{err: NewAppErr2("ResourceNotFoundErr")}, want: false},
		{name: "Unknown", args: args{err: NewAppErr2("hoge")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBadRequest(tt.args.err); got != tt.want {
				t.Errorf("IsBadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnauthorized(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ValidationErr", args: args{err: NewAppErr2("ValidationErr")}, want: false},
		{name: "WrappedValidationErr", args: args{err: NewAppErr2("WrappedValidationErr")}, want: false},
		{name: "DuplicatedIdErr", args: args{err: NewAppErr2("DuplicatedIdErr")}, want: false},
		{name: "UnauthorizedErr", args: args{err: NewAppErr2("UnauthorizedErr")}, want: true},
		{name: "ResourceNotFoundErr", args: args{err: NewAppErr2("ResourceNotFoundErr")}, want: false},
		{name: "Unknown", args: args{err: NewAppErr2("hoge")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnauthorized(tt.args.err); got != tt.want {
				t.Errorf("IsUnauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceNotFoundErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "ValidationErr", args: args{err: NewAppErr2("ValidationErr")}, want: false},
		{name: "WrappedValidationErr", args: args{err: NewAppErr2("WrappedValidationErr")}, want: false},
		{name: "DuplicatedIdErr", args: args{err: NewAppErr2("DuplicatedIdErr")}, want: false},
		{name: "UnauthorizedErr", args: args{err: NewAppErr2("UnauthorizedErr")}, want: false},
		{name: "ResourceNotFoundErr", args: args{err: NewAppErr2("ResourceNotFoundErr")}, want: true},
		{name: "Unknown", args: args{err: NewAppErr2("hoge")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotFound(tt.args.err); got != tt.want {
				t.Errorf("IsNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

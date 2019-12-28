package main

import (
	"error-handling-go/app_error"
	"error-handling-go/handler"
	"fmt"
)

func main() {
	err := app_error.NewAppErr("foo")
	if hoge, ok := err.(*app_error.HogeError); ok {
		fmt.Println(hoge)
	}

	if foo, ok := err.(*app_error.FooError); ok {
		fmt.Println(foo)
	}

	if bar, ok := err.(*app_error.BarError); ok {
		fmt.Println(bar)
	}

	if unknown, ok := err.(*app_error.UnknownError); ok {
		fmt.Println(unknown)
	}

	err = app_error.NewAppErr2("ValidationErr")
	fmt.Println(handler.HandleErr(err))

	err = app_error.NewAppErr2("WrappedValidationErr")
	fmt.Println(handler.HandleErr(err))

	err = app_error.NewAppErr2("DuplicatedIdErr")
	fmt.Println(handler.HandleErr(err))

	err = app_error.NewAppErr2("UnauthorizedErr")
	fmt.Println(handler.HandleErr(err))

	err = app_error.NewAppErr2("ResourceNotFoundErr")
	fmt.Println(handler.HandleErr(err))

	err = app_error.NewAppErr2("hoge")
	fmt.Println(handler.HandleErr(err))
}

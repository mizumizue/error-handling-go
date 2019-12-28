package app_error

type HogeError struct {
	s string
}

type FooError struct {
	s string
}

type BarError struct {
	s string
}

type UnknownError struct {
	s string
}

func (e HogeError) Error() string {
	return e.s
}

func (e FooError) Error() string {
	return e.s
}

func (e BarError) Error() string {
	return e.s
}

func (e UnknownError) Error() string {
	return e.s
}

func NewAppErr(typeS string) error {
	if typeS == "hoge" {
		return HogeError{"hoge error occurred"}
	}
	if typeS == "foo" {
		return FooError{"foo error occurred"}
	}
	if typeS == "bar" {
		return BarError{"bar error occurred"}
	}
	return UnknownError{"unknown error occurred"}
}

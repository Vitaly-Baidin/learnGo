package main

import "fmt"

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func (a *AppError) Unwrap() error {
	return a.Err
}

func (a *AppError) Error() string {
	return a.Err.Error()
}

func main() {
	err := m()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func m() error {
	return &AppError{
		Err:    fmt.Errorf("my error"),
		Custom: "value",
		Field:  404,
	}
}

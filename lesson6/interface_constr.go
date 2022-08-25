package main

import "fmt"

//var _ User = user{}

type user struct {
	Name, Address, Phone string
}

func (u *user) ChangeName(name string) {
	u.Name = name
}

func (u *user) ChangeAddress(address string) {
	u.Address = address
}

type User interface {
	ChangeName(name string)
	ChangeAddress(address string)
}

func NewUser(name, address, phone string) User {
	u := user{
		Name:    name,
		Address: address,
		Phone:   phone,
	}

	return &u
}

func main() {
	user := NewUser("name", "address", "11111")
	fmt.Println(user)
}

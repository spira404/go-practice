package main

import "fmt"

type Person struct {
	Id		int
	Name	string
	Address	string
}

type Account struct {
	Id		int
	Name	string
	Cleaner	func(string) string
	Owner	Person
}

func main() {

	// there missing fields being filled by default
	var acc Account = Account{
		Id:		1,
		Name:	"spira",
	}

	// there you neeed to fill all fields
	acc.Owner = Person{2, "spira404", "Home"}

	fmt.Printf("%#v", acc)
	// you can "inherit" one struct to another.

	// If child struct contain same field with parent
	// you need explicitly point query as Parent.Child.Field.

	// If child has field that parent hasn't, then Parent.Field
}



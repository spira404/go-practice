package main

import "fmt"

func main() {

	const pi = 3.14

	const (
		hello = "Привет"
		e     = 2.718
	)

	const (
		zero = iota // autoincrement func, eq 0
		_           // blank var, indent for iota
		two         // must be 2 bc of iota
	)

	fmt.Println(pi, hello, e, zero)
	fmt.Println(two)

	const (
		// untyped const
		// will convert (!) if needed
		year = 2017

		// typed const
		yearTyped int = 2017
	)

	var month int32 = 13

	fmt.Println(yearTyped, month+year)
}

package main

import "fmt"

func main() {

	var num0 int
	// will be 0 by defauld

	var num1 int = 1

	var num2 = 20

	num3 := 30

	num3++
	// but there is no ++num3

	userIndex := 10
	// camelCase

	var weight, height int = 10, 20
	// several vars declare

	weight, height = 11, 21
	// several vars assign

	weight, age := 12, 22
	// short vars declare+assign
	// at least one of vars must be new
	fmt.Println(num0, num1, num2, num3, userIndex, weight, height, age)
}

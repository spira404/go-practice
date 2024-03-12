package main

import "fmt"

func main() {

	a := 2
	b := &a
	*b = 4  // a = 3
	c := &a // new pointer (equal to old)

	fmt.Println(a, b, c)

	// pointer to type
	d := new(int)
	*d = 12
	*c = *d

	fmt.Println(d)

}

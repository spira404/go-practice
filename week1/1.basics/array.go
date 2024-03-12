package main

import "fmt"

func main() {

	var a1 [3]int // [0, 0, 0]
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)
	
	// can use only const to declare array size
	const size = 2
	var a2[2 * size]bool
	fmt.Println("a2", a2)

	// size depends from number of elems
	a3 := [...]int{1, 2, 3}
	fmt.Println("a3", a3)
}


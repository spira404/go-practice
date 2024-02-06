package main

import "fmt"

func main() {

	str := "name"
	switch str {
	case "name":
		fallthrough
	case "test", "lastName":
		fmt.Println("test")
	default:
		fmt.Println("not found")
	}

	// also you can break cycle inside other cycle
	// break <nameOfCycle>

	for testIdx, testRune := range str {
		fmt.Printf("%#U at index %d\n", testRune, testIdx)
	}
}

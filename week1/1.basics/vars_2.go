package main

import "fmt"

func main() {

	// platform dependent int 32/64
	var i int = 10

	var autoInt = -10

	// int8/16/32/64
	var bigInt int64 = 1<<32 - 1

	var unsignedInt uint = 100500

	// uint8/16/32/64
	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// float32/64
	var pi float32 = 3.141
	var e = 2.718
	goldenRatio := 1.618

	fmt.Println(pi, e, goldenRatio)

	// boolj
	var isOk bool = true

	fmt.Println(isOk)

	// complex64/128
	var c complex128 = -1.1 + 7.12i
	fmt.Println(c)
}

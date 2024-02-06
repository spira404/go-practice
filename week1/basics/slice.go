package main

import "fmt"

func main() {

	var buf0 []int
	buf1 := []int{1, 2, 3, 4, 5}
	buf2 := make([]int, 5, 10) // len=5 cap=10
	// capacity use to prevent too big memory allocation 
	// if the slice will stuck in the len
	fmt.Println(buf0, buf1, buf2)


	buf := append(buf1, buf2...)
	var bufLen, bufCap int = len(buf), cap(buf)

	fmt.Println(buf, bufLen, bufCap)


	newBuf := buf[:] // WRONG
	// if you change newBuf, buf will change too

	var emptyBuf []int
	copied := copy(emptyBuf, buf) // WRONG
	// mismatch of len & cap of slices lead to errors

	newBuf = make([]int, len(buf), cap(buf))
	copy(newBuf, buf) // RIGHT

	fmt.Println(newBuf, copied)
}

	 

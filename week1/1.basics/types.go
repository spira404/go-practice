package main

import "fmt"

type UserID int

func main() {
	idx := 1
	var uid UserID = 42

	// there is no type casting
	// must cast it by yourself
	myID := UserID(idx)

	fmt.Println(uid, myID)
}

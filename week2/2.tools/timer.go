package main

import(
	"fmt"
	"time"
)

func main(){
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("timer.C timeout happend")

	// short variant below (less flexible variant)
	// cant be collected by GC until it'll be finished
	case <-time.After(time.Minute):
		fmt.Println("time.After timeout happend")

	}
}

package main

import "fmt"

func main() {
	ch1 := make(chan int, 1) // 1 mean one item in buffer before lock

	go func(in chan int){
		val := <- in
		fmt.Println("Go: get from chan", val)
		fmt.Println("Go: after read from chan")
	}(ch1)

	ch1 <- 23
	ch1 <- 24

	fmt.Println("Main after put to chan")
	fmt.Scanln() // to prevent non execution from goroutines
}

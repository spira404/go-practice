package main

import "fmt"

func basicFunc(arg int) int {
	return arg + 1
}

func severalArgs(arg1, arg2 int, arg3 string) int {
	fmt.Println(arg3)
	return arg1 + arg2
}

func currentOut() (out int) {
	out = 3
	return 
	// return 2 
	// then it will return the return value (2)
}

func variatic(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	fmt.Println(currentOut())

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(variatic(nums...))
}

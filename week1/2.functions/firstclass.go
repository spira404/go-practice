package main

import "fmt"

// recover of panic
func deferTest(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened:", err)
		}
	}()
	fmt.Println("some work")
	panic("aaaaaa")
	return
}


func main() {

	deferTest()
	// lambda func
	func(in string){
		fmt.Println(in)
	}("test")

	// define func type
	type strFunc func(string)

	// closure function
	prefixer := func(prefix string) strFunc {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}
	successLogger := prefixer("SUCCESS")
	successLogger("expected behaviour")


	// if you call func in defer, it's args exec at
	// the defer define
	// to prevent this, define defer with lambda with your func exec

}

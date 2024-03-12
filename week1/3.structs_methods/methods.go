package main

import "fmt"

type Person struct {
	Id		int
	Name	string
}

// method - a function for current type
func (p *Person) SetName(name string){
	p.Name = name
}

func main() {
	pers := Person{1, "Spira"}
	pers.SetName("Spira404")
	fmt.Println(pers)

	// methods are inheritable from children (with lower priority than parent methods)

}

package main

import "fmt"

type T struct {
	name string
}

// to create a method
// you need a receiver of any type
func (t T) printName() string {
	return t.name
}

func main() {
	t := T{"Raji"}
	g := T{"g"}
	val := t.printName()
	valg := g.printName()
	fmt.Println(val, valg)
}

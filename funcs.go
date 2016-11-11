package main

// you can take the address of a function!
// also, you can change the implementation of a function under the covers!

import (
	"log"
)

/////
// func definitions
/////
type Func func(int, int) int

func newFunc(fn Func) *Func {
	return &fn
}

var fn *Func

/////
// func definitions
/////
func execute(i, j int) int {
	defer func() {
		fn = newFunc(subtract)
	}()
	return (*fn)(1, 2)
}

func add(i, j int) int      { return i + j }
func subtract(i, j int) int { return i - j }

func main() {
	fn = newFunc(add)
	log.Println(execute(1, 2))
	log.Println(execute(1, 2))
	fn = newFunc(add)
	log.Println(execute(1, 2))
}

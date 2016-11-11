package main

// you can take the address of a function!
// also, you can change the implementation of a function under the covers!

import (
	"log"
	"sync"
)

func add(i, j int) int {
	return i + j
}

func subtract(i, j int) int {
	return i - j
}

type Func func(int, int) int

func execute(fnCh <-chan *Func, contCh <-chan struct{}) int {
	fn := <-fnCh
	<-contCh
	return (*fn)(1, 2)
}

func main() {
	var wg sync.WaitGroup
	fnCh := make(chan *Func)
	contCh := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println(execute(fnCh, contCh))
	}()
	fn := Func(add)
	fnCh <- &fn
	fn = subtract
	contCh <- struct{}{}
	wg.Wait()
}

package main

import (
	"log"
)

func add(iCh <-chan *int, contCh <-chan struct{}) int {
	i := <-iCh
	<-contCh
	return *i
}

func main() {

	iCh := make(chan *int)
	contCh := make(chan struct{})

	go func() {
		i := new(int)
		*i = 1
		iCh <- i
		(*i)++
		contCh <- struct{}{}
	}()

	log.Println(add(iCh, contCh))
}

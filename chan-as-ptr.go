package main

import (
	"log"
)

func recvAndAdd(chPtr **chan int) {
	for {
		i := <-(**chPtr)
		log.Println(i + 1)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	chPtr := &ch1
	go recvAndAdd(&chPtr)

	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	*chPtr = ch2
	for i := 0; i < 10; i++ {
		ch2 <- i
	}

}

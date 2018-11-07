package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go consumer(ch)

	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("send: ", i)
	}


	//go product(ch)

	time.Sleep(1000* time.Second)

}
func consumer(ch <-chan int) {

	for {
		value := <-ch;
		fmt.Println("receive:  ", value)
	}

}
func product(ch chan<- int) {

	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("send: ", i)
	}

}

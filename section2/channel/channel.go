package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Producer function that generates random numbers and sends them to the channel
func producer(ch chan<- int) {
	for {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(time.Millisecond * 500) // Simulate some delay
	}
}

// Consumer function that receives numbers from the channel, calculates their squares, and prints them
func consumer(ch <-chan int) {
	for num := range ch {
		square := num * num
		fmt.Printf("Received %d, its square is %d\n", num, square)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	// producer and consumer as goroutines
	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second * 10)
}

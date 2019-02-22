package main

import (
	"fmt"
	"time"

	"github.com/patterjm/go-sandbox/chans-sandbox/chans"
)

func main() {

	fmt.Println("Adding to and reading one message from a channel")
	time.Sleep(750 * time.Millisecond)
	runChansGetPing()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Add and read two messages in a buffered channel")
	time.Sleep(750 * time.Millisecond)
	runChansGetBufferedPings()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Spawning 10 worker threads and using a waitgroup to coordinate closing the channel")
	time.Sleep(750 * time.Millisecond)
	runChansGetPingNTimes(10)
}

func runChansGetPing() {
	messages := make(chan string)
	go chans.GetPing(messages)
	msg := <-messages
	fmt.Println(msg)
}

func runChansGetBufferedPings() {
	messages := make(chan string, 2)
	go chans.GetPing(messages)
	go chans.GetPong(messages)
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}

func runChansGetPingNTimes(num int) {
	messages := make(chan string)
	go chans.GetPingNTimes(messages, num)
	for msg := range messages {
		fmt.Println(msg)
	}
}

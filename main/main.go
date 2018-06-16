package main

import (
	"fmt"

	"github.com/patterjm/go-sandbox/chans"
)

func main() {
	fmt.Println("Hello world")
	runChansGetPing()
	runChansGetPingNTimes(10)
}

func runChansGetPing() {
	messages := make(chan string)
	go chans.GetPing(messages)
	msg := <-messages
	fmt.Println(msg)
}

func runChansGetPingNTimes(num int) {
	messages := make(chan string)
	go chans.GetPingNTimes(messages, num)
	for msg := range messages {
		fmt.Println(msg)
	}
}

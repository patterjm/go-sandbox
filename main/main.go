package main

import (
	"fmt"

	"github.com/patterjm/go-sandbox/chans"
)

func main() {
	fmt.Println("Hello world")
	messages := make(chan string)
	go chans.GetPing(messages)
	msg := <-messages
	fmt.Println(msg)

}

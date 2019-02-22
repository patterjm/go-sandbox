package chans

import (
	"sync"
)

//GetPing Adds string "ping" to given channel
func GetPing(messages chan (string)) {
	messages <- "ping"
}

//GetPingNTimes Adds string "ping" to given channel by creating n go routines and adding 1 message per goroutine
func GetPingNTimes(messages chan (string), threads int) {
	var wg sync.WaitGroup
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go getPingWithWG(messages, &wg)
	}
	wg.Wait()
	close(messages)
}

func getPingWithWG(messages chan (string), waitGroup *sync.WaitGroup) {
	GetPing(messages)
	waitGroup.Done()
}

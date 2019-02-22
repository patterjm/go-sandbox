package chans

import (
	"fmt"
	"sync"
	"time"
)

//GetPing Adds string "ping" to given channel
func GetPing(messages chan (string)) {
	messages <- "ping"
}

//GetPong Adds string "pong" to given channel
func GetPong(messages chan (string)) {
	messages <- "pong"
}

//GetPingNTimes Adds string "ping" to given channel by creating n go routines and adding 1 message per goroutine
func GetPingNTimes(messages chan (string), threads int) {
	var wg sync.WaitGroup
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go getPingWithWG(messages, i, &wg)
	}
	wg.Wait()
	close(messages)
}

func getPingWithWG(messages chan (string), threadNum int, waitGroup *sync.WaitGroup) {
	time.Sleep(3000 * time.Millisecond)
	getPingWithThreadNum(messages, threadNum)
	waitGroup.Done()
}

func getPingWithThreadNum(messages chan (string), threadNum int) {
	messages <- fmt.Sprintf("Wait group chan %s thread: %d", "ping", threadNum)
}

package syncStepGoroutines

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Msg struct {
	val         string
	syncChannel chan bool
}

func someFunc(ch chan Msg, goroutineName string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		msg := <-ch
		// do some operation with the message
		// received
		fmt.Println(goroutineName, msg.val, i)
		// blocked; wait for go-ahead
		msg.syncChannel <- true
	}
}

func Synched() {
	// used to send messages to goroutines
	msg1Ch := make(chan Msg)
	msg2ch := make(chan Msg)
	// shared between all messages
	commonSyncChannel := make(chan bool)
	wg.Add(2)
	go someFunc(msg1Ch, "Goroutine 1, ")
	go someFunc(msg2ch, "Goroutine 2, ")
	for i := 0; i < 5; i++ {
		// commonSyncChannel is shared between all messages
		msg1Ch <- Msg{fmt.Sprintf("msg%d", i), commonSyncChannel}
		msg2ch <- Msg{fmt.Sprintf("msg%d", i), commonSyncChannel}
		// some operation on current routine
		time.Sleep(time.Second * 2)
		// give the go-ahead for both goroutines
		<-commonSyncChannel
		<-commonSyncChannel
	}
	wg.Wait()
}

package quitindicator

import (
	"fmt"
	"time"
)

func someFunc(ch chan int, quit chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, <-ch)
	}
	// indicate that the goroutine wants to quit
	quit <- true
	if <-quit {
		fmt.Println("Goroutine Quit Acknowledged")
	}
}

func QuitIndicatorExample() {
	ch := make(chan int)
	quitSignalCh := make(chan bool)
	defer close(ch)
	defer close(quitSignalCh)
	go someFunc(ch, quitSignalCh)
	for i := 100; ; i++ {
		select {
		case <-quitSignalCh:
			fmt.Println("Goroutine has indicated that it wants to quit")
			// perform some final operation before total quit
			// signal the Goroutine that all is done and it can stop now for sure
			quitSignalCh <- true
			time.Sleep(time.Second)
			return
		case ch <- i:
			// do nothing
		}
	}
}

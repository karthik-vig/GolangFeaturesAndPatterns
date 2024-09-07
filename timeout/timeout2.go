package timeout

import (
	"fmt"
	"time"
)

func someFunc2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 800)
	}
}

func TimeoutType2() {
	ch := make(chan int)
	go someFunc2(ch)
	timeoutCh := time.After(time.Second * 5)
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-timeoutCh:
			fmt.Println("The time receive values is over...")
			return
		}
	}
}

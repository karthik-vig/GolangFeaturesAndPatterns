package timeout

import ( 
	"time"
	"fmt"
)

func someFunc(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 900)
		if i == 2 {
			time.Sleep(time.Millisecond * 1200)
		}
	}
}

func TimeoutType1() {
	ch := make(chan int)
	go someFunc(ch)
	for {
		select {
			case val := <-ch:
				fmt.Println(val)
			case <-time.After(time.Second * 1):
				fmt.Println("It took too long to respond; quiting...")
				return
		}
	}
}
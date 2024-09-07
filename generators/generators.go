package generators

import "time"

func GeneratorEx(startValue int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := startValue; i < startValue+10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	return ch
}

package multiplexing

import "sync"

// input channels are receive only channels
// we return a new receive only output channel
func Multiplex(ch1, ch2 <-chan int) <-chan int {
	commonCh := make(chan int)
	getValFunc := func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			val, ok := <-ch
			if !ok {
				break
			}
			commonCh <- val
		}
	}
	go func() {
		defer close(commonCh)
		var wg sync.WaitGroup
		wg.Add(2)
		go getValFunc(ch1, &wg)
		go getValFunc(ch2, &wg)
		wg.Wait()
	}()
	return commonCh
}

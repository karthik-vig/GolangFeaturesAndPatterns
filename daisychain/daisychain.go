package daisychain

import "fmt"

func someFunc(left, right chan int) {
	left <- 2 + <-right
}

func DaisyChainExample() {
	firstRightCh := make(chan int)
	rightCh := firstRightCh
	var leftCh chan int
	for i := 0; i < 100000; i++ {
		leftCh = make(chan int)
		go someFunc(leftCh, rightCh)
		rightCh = leftCh
	}
	firstRightCh <- 1
	fmt.Println(<-leftCh)
}

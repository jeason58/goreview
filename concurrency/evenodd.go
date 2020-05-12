package concurrency

import (
	"fmt"
	"sync"
)

func PrintEvenOddByMultiThread() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			select {
			case n := <-evenCh:
				fmt.Printf("%d\t", n)
				oddCh <- n+1
			}
		}
	}()

	go func() {
		for {
			select {
			case n := <-oddCh:
				fmt.Printf("%d\t", n)
				evenCh <- n+1
			}
		}
	}()

	evenCh <- 0
	wg.Wait()
}

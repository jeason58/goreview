package algorithm

import (
	"fmt"
	"time"
)

func Fibonacci() {
	go spinner(100 * time.Millisecond)
	const n = 10
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	fn1 := 1
	fn2 := 0
	fn := 0
	for i := 2; i <= x; i++ {
		fn = fn1 + fn2
		fn2 = fn1
		fn1 = fn
	}

	time.Sleep(time.Second * 2)

	return fn
}


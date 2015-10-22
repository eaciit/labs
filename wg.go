package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func fib(d int) int {
	r := 0
	for i := 1; i <= d; i++ {
		r += i
	}
	return r
}

func main() {
	wg := new(sync.WaitGroup)
	completed := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(d int, c *int) {
			fib := fib(d)
			fmt.Printf("Writing: %d fib: %d\n", d, fib)
			ifloat := float64(d)
			if math.Mod(ifloat, 100) == 0 && d != 0 {
				time.Sleep(time.Duration(fib) * time.Millisecond)
				fmt.Printf("Sending data: %d \n", d)
			}
			*c = *c + 1
			wg.Done()
		}(i, &completed)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("All seed has been sent")
	wg.Wait()
	fmt.Printf("Done, completed %d\n", completed)
}

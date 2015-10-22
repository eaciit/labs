package main

import (
	"fmt"
	//"github.com/eaciit/toolkit"
	"sync"
	"time"
)

type FibObj struct {
	Seed int
	Fib  int
}

func fib(in interface{}) interface{} {
	d := in.(int)
	r := 0
	for i := 1; i <= d; i++ {
		r += i
	}
	//randtime := toolkit.RandInt(1)
	//time.Sleep(time.Duration(randtime) * time.Millisecond)
	time.Sleep(10 * time.Millisecond)
	return r
}

func main() {
	fmt.Println("Init ...")

	workerCount := 10
	jobCount := 100
	preparedJob := 0
	processedJob := 0
	completedJob := 0
	wg := new(sync.WaitGroup)

	allKeysSent := false
	running := true

	keys := make(chan interface{})
	results := make(chan interface{})
	keySentChannel := make(chan bool)

	for widx := 0; widx < workerCount; widx++ {
		go func(ks chan interface{}, rs chan interface{}, wg *sync.WaitGroup,
			processedJob *int,
			fn func(in interface{}) interface{}) {
			var ki int
			for !allKeysSent {
				for k := range ks {
					*processedJob = *processedJob + 1
					ki = k.(int)
					rs <- FibObj{ki, fn(ki).(int)}
					wg.Done()
				}
			}
		}(keys, results, wg, &processedJob, fib)
	}

	go func() {
		for !allKeysSent {
			select {
			case <-keySentChannel:
				allKeysSent = true
			}
		}
	}()

	go func(completedJob *int, running *bool) {
		for *running == true {
			select {
			case r := <-results:
				*completedJob = *completedJob + 1
				fmt.Printf("Prepared: %d, Processed: %d, completed %d | Fibo results: %v \n",
					preparedJob, processedJob, *completedJob, r.(FibObj))
			}
		}
	}(&completedJob, &running)

	go func(wg *sync.WaitGroup, preparedJob *int) {
		for ki := 0; ki < jobCount; ki++ {
			wg.Add(1)
			*preparedJob = *preparedJob + 1
			//preparedJob = preparedJob + 1
			keys <- ki
		}
		close(keys)
		keySentChannel <- true
	}(wg, &preparedJob)

	for !allKeysSent {
		time.Sleep(1 * time.Millisecond)
	}
	wg.Wait()
	running = false

	//time.Sleep(5 * time.Second)
	fmt.Printf("Done, completed %d jobs \n", completedJob)
}

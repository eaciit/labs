package main
//test fork
import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	in := make(chan string)
	scanDone := make(chan int)
	//out := make(chan string)

	go func() { in <- strconv.Itoa(10) }()
	fmt.Println(<-in)

	i := 0
	for i < 100 {
		go func(i int) {
			s := strconv.Itoa(i + 1)
			fmt.Println("Sending " + s)
			in <- s
		}(i)
		i++
	}
	go func() {
		fmt.Printf("Data scanned send: %d \n", i)
		scanDone <- i
	}()

	finish := false
	scanHasBeenCompleted := false
	dataProcessed := 0
	dataRcvd := 0
	for !finish {
		select {
		case dataProcessed = <-scanDone:
			fmt.Printf("Data scanned %d \n", dataProcessed)
			scanHasBeenCompleted = true
		case s := <-in:
			dataRcvd++
			fmt.Printf("Receiving %d data, processed %d. Last value %s \n", dataRcvd, dataProcessed, s)
		default:
			if dataRcvd == dataProcessed && scanHasBeenCompleted {
				finish = true
			}
			//-- do nothing
		case <-time.After(5 * time.Second):
			finish = true
		}

	}

	fmt.Println("Done")
}

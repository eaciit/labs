package main

import (
	"flag"
	"fmt"
	_ "github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
	"github.com/eaciit/toolkit"
)

var (
	f = flow.New()
)

func init() {

}

func main() {
	flag.Parse()

	f.Source(func(out chan string) {
		for i := 1; i <= 1000; i++ {
			txt := createRandomString(toolkit.RandInt(23) + 10)
			fmt.Printf("Data %d is %s \n", i, txt)
			out <- txt
		}
	}, 3).Map(func(s string) (int, int) {
		return len(s), 1
	}).Partition(5).ReduceByKey(func(x, y int) int {
		return x + y
	}).Sort(nil).Map(func(k int, v int) {
		fmt.Printf("Number of data with %d chars are %d \n", k, v)
	})

	flow.Ready()
	f.Run()
}

func createRandomString(randomLength int) string {
	randomTxt := ""
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890!@#$%_-"
	for x := 0; x < randomLength; x++ {
		ic := toolkit.RandInt(len(chars) - 1)
		c := chars[ic]
		randomTxt += string(c)
	}
	return randomTxt
}

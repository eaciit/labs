package main

import (
	"fmt"
	tk "github.com/eaciit/toolkit"
	"time"
)

func todate(source string) time.Time {
	d, e := time.Parse("1/2/06", source)
	if e == nil {
		return d
	}
	return time.Now()
}

func main() {
	s1 := "9/30/12"
	s2 := "1/10/14"

	fmt.Printf("Date %v \n", tk.MakeDate("1/2/06", s1))
	fmt.Printf("Date %v \n", tk.MakeDate("1/2/06", s2))
}

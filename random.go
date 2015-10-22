package main

import (
	"fmt"
	tk "github.com/eaciit/toolkit"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Print(tk.RandInt(100), ",")
	}
	fmt.Println()
}

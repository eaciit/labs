package main

import (
	"fmt"
	"github.com/eaciit/toolkit"
	"os"
)

func main() {
	o, e := toolkit.RunCommand("go", "env")
	if e == nil {
		fmt.Printf("Command %s result \n%v\n", "ls -al", o)
	} else {
		fmt.Printf("Unable to run command: %s \n", e.Error())
	}
}

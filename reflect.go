package main

import (
	"fmt"
	tk "github.com/eaciit/toolkit"
	"reflect"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

func main() {
	//emp := new(Employee)
	emps := []Employee{}
	fillArray(&emps)
	fmt.Printf("Value of empls:\n%s\n", tk.JsonString(emps))
}

func fillArray(o interface{}) {
	//rt := reflect.TypeOf(o)
	isPointer := false
	rv1 := reflect.ValueOf(o)
	var rv2 reflect.Value

	if rv1.Kind() == reflect.Ptr {
		isPointer = true
		rv2 = reflect.Indirect(rv1)
	} else {
		rv2 = rv1
	}

	if rv2.Kind() == reflect.Slice {
		//fmt.Println("Add new element")
		rv2 = reflect.Append(rv2, reflect.ValueOf(Employee{"EC0003", "Arief Darmawan", 30}))
		if isPointer {
			rv1.Elem().Set(rv2)

		}
	} else {
		fmt.Printf("Kind: %s\n", rv2.Kind().String())
	}
}

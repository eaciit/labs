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
	//var emps interface{}
	emps := tk.MakeSlice(&Employee{}).([]*Employee)
	fillArray(&emps)
	fmt.Printf("Value of empls:\n%s\n", tk.JsonString(emps))
}

func fillArray(o interface{}) {
	rv1 := reflect.ValueOf(o)
	var rv2 reflect.Value

	if rv1.Kind() == reflect.Ptr {
		rv2 = reflect.Indirect(rv1)
	} else {
		fmt.Println("Object passed is not a ptr")
		return
	}

	if rv2.Kind() == reflect.Slice {
		tslice := rv2.Type().Elem()
		var newEmp reflect.Value
		if string(tslice.String()[0]) == "*" {
			newEmp = reflect.New(tslice.Elem())
		} else {
			newEmp = reflect.Indirect(reflect.New(tslice.Elem()))
		}
		rv2 = reflect.Append(rv2, newEmp)
		rv1.Elem().Set(rv2)
	} else {
		fmt.Printf("Kind: %s\n", rv2.Kind().String())
	}
}

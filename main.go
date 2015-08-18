package main

import (
	"fmt"
	. "github.com/eaciit/toolkit"
)

func main() {
	m := M{"data": M{"_id": 1, "title": "Variable 1"}, "count": 20, "obj": struct {
		Id    interface{}
		Title string
	}{300, "Object 20"}}
	fmt.Printf("ID of object 1 is %v \n", Id(m["obj"]))
	fmt.Printf("Title value of object 1.obj.Title is %v \n", Value(Value(m, "obj", nil), "Title", ""))
	//fmt.Printf("Count is %v \n", m["count"])
}

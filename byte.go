package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/eaciit/toolkit"
)

func main() {
	type Obj struct {
		AA string
		BB string
	}
	data := []Obj{
		{"A", "B"},
		{"C", "D"},
		{"E", "F"}}

	b := toolkit.GetEncodeByte(data)
	fmt.Printf("Data is : %v\nString: %v\nLen: %d\n", b, string(b), len(b))
	var decoded Obj
	toolkit.DecodeByte(b, &decoded)
	fmt.Printf("Decoded: %v \n\n", decoded)

	//b = []byte(fmt.Sprint(data))
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	b = buf.Bytes()
	fmt.Printf("Data is : %v\nString: %v\nLen: %d\n", b, string(b), len(b))
	eread := binary.Read(buf, binary.LittleEndian, &decoded)
	if eread != nil {
		fmt.Printf("Unable to decode. %s \n", eread.Error())
	} else {
		fmt.Printf("Decoded: %v \n\n", decoded)
	}
}

package main

import (
	. "github.com/eaciit/toolkit"
	"time"
)

type Obj struct {
	ID, Name string
	Age      int
	Enable   bool
	Created  time.Time
}

func main() {
	o := new(Obj)
	o.ID = "eaciit"
	o.Name = "EACIIT"
	o.Age = RandInt(40)
	o.Enable = true
	o.Created = time.Now()

	Printf("Original value: %s \n", JsonString(o))
	// -- SerDe to from Obj to ms
	ms := M{}
	e := FromBytes(ToBytes(o, ""), "", &ms)
	if e != nil {
		Printf("Serde using bytes fail: %s\n", e.Error())
		return
	}
	delete(ms, "Age")
	Printf("Map[string]interface{} Value: %s\n", JsonString(ms))

	// -- Now change some value on ms
	ms["Name"] = "PT Wiyasa Teknologi Nusantara"
	ms["Created"] = time.Date(2013, 4, 1, 0, 0, 0, 0, time.Now().Location())

	// -- Lets serde back from ms to o
	e = FromBytes(ToBytes(ms, ""), "", o)
	Printf("Object value after serde and be changed on M: %s \n", JsonString(o))
}

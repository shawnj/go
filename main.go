package main

import (
	"fmt"

	"myapp/newmessage"
)

// NewMessage - is a struct

func main() {

	var s = newmessage.NewMessage{Name: "Joe", Message: "Hello newMess"}

	fmt.Println(s.NewMess())

	testOne()
}

func testOne() {
	//msg := "hello there"
	a, b, c := 1, 2, 3

	var aa *int

	aa = &a

	fmt.Println(*aa, a, b, c)
}

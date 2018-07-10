package main

import (
	"fmt"
	"time"

	newmessage "myapp/newmessage"
)

type GetName interface {
	DisplayName() string
	UpdateName(string)
}

type errorer func() error

func main() {

	var msgDis GetName
	msgDis = &newmessage.NewMessage{}
	msgDis.UpdateName("Joe")

	s := newmessage.NewMessage{}
	s.UpdateName("Mike")
	s.UpdateMessage("Hello There")

	fmt.Println(msgDis.DisplayName() + " " + s.DisplayMsg())

	testOne()

	go func() {
		fmt.Println("hello func")
	}()

	time.Sleep(time.Millisecond)
}

func testOne() {
	//msg := "hello there"
	a, b, c := 1, 2, 3

	// var aa *int

	aa := &a

	fmt.Println(aa, a, b, c)
}

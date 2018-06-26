package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world there")
	testOne()
}

func testOne() {
	//msg := "hello there"
	a, b, c := 1, 2, 3

	var a2 = &a

	fmt.Println(a2, a, b, c)
}

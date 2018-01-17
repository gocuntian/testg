package main

import (
	"fmt"
	"time"
)

func main() {
	a()
	fmt.Println("normally returned from main")
}

// Inside A
// Inside B
// recovered: oh! B panicked
// normally returned from main

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	defer recovery()
	fmt.Println("Inside B")
	panic("oh! B panicked")
}
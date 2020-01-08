package main

import (
	"fmt"
	"time"
)

type greeting string
type n int

func (g greeting) Greet() {
	for i := 0; i < 10; i++ {
		timer1 := time.NewTimer(2 * time.Second)
		<-timer1.C
		fmt.Println("value ", i)

		sum := 0
		sum += i * 100
		fmt.Println(sum)
	}
	
}

// exported
var Greeter greeting

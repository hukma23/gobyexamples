package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 10 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

package main

import (
	"fmt"
	"time"
)

func main6() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- idx:
				fmt.Println(idx, "Send result")
			case <-done:
				fmt.Println(idx, "Exiting")
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Result: ", <-ch)
	close(done)
	time.Sleep(1 * time.Second)
}

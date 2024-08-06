package main

import "runtime"

func main18() {
	done := false

	go func() {
		done = true
	}()

	for !done {
		//println("not done !")
		runtime.Gosched()
	}

	println("done !")
}

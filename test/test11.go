package main

import "fmt"

func main11() {
	defer func() {
		doRecover()
	}()
	panic("not good")
}

func doRecover() {

	defer func() {
		fmt.Println("recobered: ", recover())
	}()
	panic("woce")
}

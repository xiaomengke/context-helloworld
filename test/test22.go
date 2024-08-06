package main

import "fmt"

type Direction int

var p2 *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p2)
}

func main22() {
	//var err error
	//p2, err = foo()
	p2, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p2)
}

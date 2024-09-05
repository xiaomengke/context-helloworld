package main

import (
	"fmt"
)

type User struct{}
type User1 User
type User2 = User //别名

func (i User1) m1() {
	fmt.Println("m1")
}
func (i User) m2() {
	fmt.Println("m2")
}

func main27() {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()
}

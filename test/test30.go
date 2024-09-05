package main

import (
	"fmt"
)

type N int

func (n N) test() {
	fmt.Println(n)
}

func mainx() {
	var n N = 10

	fmt.Println(n)

	n++
	N.test(n)

	n++
	(*N).test(&n)
}
func mainc() {
	s := make([]int, 9, 9)
	fmt.Println(len(s))
	s2 := s[4:6]
	s2[1] = 9
	fmt.Println(s)
	fmt.Println(s2)

	s2 = append(s2, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(len(s2))
	fmt.Println(s)
	fmt.Println(s2)

}

func main30() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	_ = y == y
	//unix.Getegid()
	//sync.WaitGroup{}
	//reflect.SliceHeader{}
	//map[]
	//uintptr()
	//reflect.ValueOf()
}

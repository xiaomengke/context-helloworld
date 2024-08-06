package main

import "fmt"

// 1.
func main33() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 2.
func main3() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}
func funcMui(x, y int) (sum int, err error) {
	//sum =x+y
	return x + y, nil
}

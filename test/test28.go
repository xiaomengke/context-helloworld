package main

import "fmt"

func main28() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c)

	var s []int
	s = append(s, 1)

	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1

	x1 := []int{ // 多行
		1,
		2,
	}
	x1 = x1

	y1 := []int{3, 4} // 一行 no error
	y1 = y1

	var aa byte = 0x11
	fmt.Println(aa)
	var bb uint8 = aa
	var cc uint8 = aa + bb
	test(cc)

}
func test(x byte) {
	fmt.Println(x)
}

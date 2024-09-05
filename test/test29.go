package main

import "fmt"

var ch chan int = make(chan int, 9)

func main29() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x) // print ?
	ch <- 1
	<-ch
	close(ch)

	var d int = 3
	var a int8 = 3
	var b uint8 = 3
	var c int8 = -3
	//有符号取反 = 加1取负
	//无符号取反 = 按位取反
	fmt.Printf("^%b=%b %d\n", d, ^d, ^d)
	fmt.Printf("^%b=%b %d\n", a, ^a, ^a) // ^11=-100 -4
	fmt.Printf("^%b=%b %d\n", b, ^b, ^b) // ^11=11111100 252
	fmt.Printf("^%b=%b %d\n", c, ^c, ^c) // ^-11=10 2
}

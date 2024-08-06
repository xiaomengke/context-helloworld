package main

import (
	"fmt"
	"strings"
)

type MyInt1 int
type MyInt2 = int //别名

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main19() {
	var i int = 0
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...) //必须加...
	fmt.Println(s1)

	//字符串拼接
	str := "avd" + "213"
	s := fmt.Sprintf("acd%d", 124)
	ss := []string{"asd", "sadax"}
	sss := strings.Join(ss, "")
	fmt.Println(str, s, sss)

	//
	var x error = nil ///nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
	var y interface{} = nil
	//var z string = nil
	_, _ = x, y
	//m:=map[int]int{1:3}
	//fmt.Println(cap(m))//cap不使用map

}

package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

type People2 interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main21() {
	var peo People2 = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}

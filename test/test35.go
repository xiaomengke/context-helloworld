package main

import (
	"fmt"
	"sync"
)

type demo struct {
}

type Instance interface {
	work() string
}

func (d *demo) work() string {
	return "1243"
}
func newDemo() *demo {
	return &demo{}
}

var dem *demo
var once sync.Once

func GetDemo() Instance {
	once.Do(func() {
		dem = newDemo()
	})
	return dem
}

func main35() {
	demoInstance := GetDemo().(*demo)
	demoInstance2 := GetDemo().(*demo)
	fmt.Println(&demoInstance)
	fmt.Println(&demoInstance2)
}

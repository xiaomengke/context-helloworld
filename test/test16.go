package main

import "fmt"

//	func First(query string, replicas []Search) Result {
//		c := make(chan Result)
//		done := make(chan struct{})
//		defer close(done)
//		replicaSearch := func(i int) {
//			select {
//			case c <- replicas[i](query):
//			case done:
//		}
//		for i := range replicas {
//			go replicaSearch(i)
//		}
//		return <-c
//	}
type data2 struct {
	name string
}

func (p *data2) print() {
	fmt.Println("name: ", p.name)
}

func main16() {
	//d1 := data2{"one"}
	//d1.print() // d1 变量可寻址，可直接调用指针 receiver 的方法

	m := map[string]*data2{ //带*
		"x": &data2{"three"}, //存地址就行了
	}
	m["x"].print()
	m["x"].name = "123"
	temp := m["x"]
	temp.print()
	temp.name = "..."
	m["x"].print()
}
